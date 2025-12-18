package git

import (
	"bytes"
	"os/exec"
)

// Commit은 주어진 메시지로 커밋을 생성합니다
func Commit(message string, noVerify bool) error {
	args := []string{"commit", "-m", message}
	
	if noVerify {
		args = append(args, "--no-verify")
	}
	
	cmd := exec.Command("git", args...)
	
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	
	if err := cmd.Run(); err != nil {
		return err
	}
	
	return nil
}

// GetLastCommitMessage는 마지막 커밋 메시지를 반환합니다
func GetLastCommitMessage() (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=%B")
	
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	if err := cmd.Run(); err != nil {
		return "", err
	}
	
	return stdout.String(), nil
}
