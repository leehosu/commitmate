package git

import (
	"bytes"
	"os/exec"
	"strings"
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
	return diff, nil
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
