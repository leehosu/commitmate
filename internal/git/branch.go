package git

import (
	"os/exec"
	"regexp"
	"strings"
)

// GetCurrentBranch returns the current git branch name
func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// ExtractJiraIssue extracts JIRA issue key from branch name
// Pattern: PROJECT-123, ABC-456, DEVOPS2-430, etc.
func ExtractJiraIssue(branchName string) string {
	// JIRA 이슈 패턴: 대문자로 시작, 숫자 포함 가능, 하이픈, 숫자
	re := regexp.MustCompile(`([A-Z]+[A-Z0-9]*-[0-9]+)`)
	match := re.FindString(branchName)
	return match
}

// FormatCommitMessage adds JIRA issue key if present in branch name
// If jiraIntegration is false, explicitly skip JIRA detection
// If true or not set, automatically detect and add JIRA issue if present
func FormatCommitMessage(message string, jiraIntegration bool) (string, error) {
	// JIRA 통합이 명시적으로 비활성화되어 있으면 원본 반환
	if !jiraIntegration {
		return message, nil
	}

	// Get current branch
	branch, err := GetCurrentBranch()
	if err != nil {
		return message, nil // 에러 시 원본 메시지 반환 (silent fail)
	}

	// main, master, develop 등 특수 브랜치는 제외
	if branch == "main" || branch == "master" || branch == "develop" || branch == "HEAD" {
		return message, nil
	}

	// Extract JIRA issue - 브랜치에 JIRA 패턴이 있는지 자동 감지
	jiraIssue := ExtractJiraIssue(branch)
	if jiraIssue == "" {
		return message, nil // JIRA 이슈 패턴이 없으면 원본 반환
	}

	// Already has JIRA issue? Skip
	if strings.Contains(message, jiraIssue) {
		return message, nil
	}

	// Add JIRA issue prefix
	return "[" + jiraIssue + "] " + message, nil
}
