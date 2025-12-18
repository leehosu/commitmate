package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const (
	// MaxDiffSize는 AI에게 전달할 최대 diff 크기 (약 10KB, ~2500 토큰)
	MaxDiffSize = 10000
)

// GetStagedDiff는 staged 변경사항의 diff를 반환합니다
func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	if err := cmd.Run(); err != nil {
		return "", err
	}
	
	diff := strings.TrimSpace(stdout.String())
	
	// Diff가 너무 크면 요약
	if len(diff) > MaxDiffSize {
		return getSummarizedDiff()
	}
	
	return diff, nil
}

// getSummarizedDiff는 큰 diff를 요약합니다
func getSummarizedDiff() (string, error) {
	// git diff --stat으로 통계 가져오기
	statCmd := exec.Command("git", "diff", "--cached", "--stat")
	var statOut bytes.Buffer
	statCmd.Stdout = &statOut
	if err := statCmd.Run(); err != nil {
		return "", err
	}
	
	// git diff --shortstat으로 요약 가져오기
	shortStatCmd := exec.Command("git", "diff", "--cached", "--shortstat")
	var shortStatOut bytes.Buffer
	shortStatCmd.Stdout = &shortStatOut
	if err := shortStatCmd.Run(); err != nil {
		return "", err
	}
	
	// 변경된 파일 목록
	files, err := GetStagedFiles()
	if err != nil {
		return "", err
	}
	
	summary := fmt.Sprintf(`[Large changeset - Summary provided]

Changed files (%d):
%s

Statistics:
%s
%s

Note: Full diff is too large. Please review the changes manually if needed.
`, len(files), strings.Join(files, "\n"), statOut.String(), shortStatOut.String())
	
	return summary, nil
}

// CheckRepository는 현재 디렉토리가 git 저장소인지 확인합니다
func CheckRepository() error {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	
	if err := cmd.Run(); err != nil {
		return err
	}
	
	return nil
}

// GetStagedFiles는 staged 파일 목록을 반환합니다
func GetStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	
	files := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(files) == 1 && files[0] == "" {
		return []string{}, nil
	}
	
	return files, nil
}
