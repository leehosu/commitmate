package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/leehosu/commitmate/internal/ai"
	"github.com/leehosu/commitmate/internal/config"
	"github.com/leehosu/commitmate/internal/git"
	"github.com/leehosu/commitmate/internal/i18n"
	"github.com/leehosu/commitmate/internal/prompt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "AI가 생성한 커밋 메시지로 커밋합니다",
	Long:  `Staged 변경사항을 분석하여 AI가 자동으로 커밋 메시지를 생성하고 커밋합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runCommit(); err != nil {
			// 에러 발생 시 UI 언어에 맞게 출력
			cfg, cfgErr := config.Load()
			if cfgErr != nil {
				cfg = config.Default()
			}
			msg := i18n.GetMessages(cfg.UILanguage)
			
			// 에러 메시지 출력 (i18n 적용)
			color.Red("❌ %s: %v", msg.ErrorInputFailed, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}

func runCommit() error {
	// 3. 설정 로드
	cfg, err := config.Load()
	if err != nil {
		// 설정 로드 실패 시에도 기본 설정으로 메시지 표시
		defaultCfg := config.Default()
		msg := i18n.GetMessages(defaultCfg.UILanguage)
		return fmt.Errorf("%s: %w", msg.ErrorLoadConfig, err)
	}

	// UI 메시지 가져오기
	msg := i18n.GetMessages(cfg.UILanguage)

	// 1. Git 저장소 확인
	color.Cyan(msg.CheckingRepository)
	if err := git.CheckRepository(); err != nil {
		return fmt.Errorf("%s: %w", msg.ErrorNotGitRepo, err)
	}

	// 2. Staged 변경사항 가져오기
	color.Cyan(msg.AnalyzingStagedChanges)
	diff, err := git.GetStagedDiff()
	if err != nil {
		return fmt.Errorf("%s: %w", msg.ErrorNoStagedChanges, err)
	}

	if diff == "" {
		return errors.New(msg.NoChanges)
	}

	// Diff 크기가 크면 경고
	if len(diff) > 5000 {
		color.Yellow(msg.LargeChangesWarning, len(diff))
	}

	// 명령줄에서 지정한 provider가 있으면 덮어쓰기
	if provider != "" {
		cfg.Provider = provider
	}

	// 4. AI 클라이언트 생성
	client, err := ai.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("%s\n\n%s", fmt.Sprintf(msg.ErrorCreateClient, err), fmt.Sprintf(msg.HintSetAPIKey, cfg.Provider))
	}

	// 5. 커밋 메시지 생성
	color.Cyan(msg.GeneratingMessage)

	// 프롬프트 생성 (CommitLanguage 사용)
	systemPrompt := prompt.GetSystemPrompt(cfg.CommitLanguage, cfg.Template)
	userPrompt := prompt.GetUserPrompt(diff)

	commitMessage, err := client.GenerateCommitMessage(systemPrompt, userPrompt)
	if err != nil {
		return fmt.Errorf("%s: %w", msg.ErrorGenerateMessage, err)
	}

	// JIRA 이슈 번호 추가
	originalMessage := commitMessage
	commitMessage, _ = git.FormatCommitMessage(commitMessage, cfg.JiraIntegration)
	
	// JIRA 이슈가 추가되었으면 사용자에게 알림
	if originalMessage != commitMessage {
		branch, _ := git.GetCurrentBranch()
		jiraIssue := git.ExtractJiraIssue(branch)
		if jiraIssue != "" {
			color.Cyan(msg.JiraIssueDetected, jiraIssue)
			color.Green(msg.JiraIssueAdded)
		}
	}

	// 6. 생성된 메시지 출력
	fmt.Println()
	color.Green(msg.MessageGenerated)
	fmt.Println()
	color.Yellow("%s", commitMessage)
	fmt.Println()

	// dry-run 모드면 여기서 종료
	if dryRun {
		color.Cyan(msg.DryRunMode)
		return nil
	}

	// 7. 사용자 확인
	for {
		promptSelect := promptui.Select{
			Label: msg.PromptLabel,
			Items: []string{
				msg.PromptYes,
				msg.PromptEdit,
				msg.PromptRegenerate,
				msg.PromptCancel,
			},
		}

		_, result, err := promptSelect.Run()
		if err != nil {
			return fmt.Errorf("%s: %w", msg.ErrorSelectFailed, err)
		}

		switch result {
		case msg.PromptYes:
			// 8. 커밋 실행
			color.Cyan(msg.Committing)
			if err := git.Commit(commitMessage, noVerify); err != nil {
				return fmt.Errorf("%s: %w", msg.ErrorCommitFailed, err)
			}
			color.Green(msg.CommitSuccess)
			return nil

		case msg.PromptEdit:
			// 수정 프롬프트
			promptEdit := promptui.Prompt{
				Label:   msg.EditPromptLabel,
				Default: commitMessage,
			}
			editedMessage, err := promptEdit.Run()
			if err != nil {
				return fmt.Errorf("%s: %w", msg.ErrorInputFailed, err)
			}

			// 수정된 메시지로 커밋
			color.Cyan(msg.Committing)
			if err := git.Commit(editedMessage, noVerify); err != nil {
				return fmt.Errorf("%s: %w", msg.ErrorCommitFailed, err)
			}
			color.Green(msg.CommitSuccess)
			return nil

		case msg.PromptRegenerate:
			// 다시 생성
			color.Cyan(msg.RegeneratingMessage)
			commitMessage, err = client.GenerateCommitMessage(systemPrompt, userPrompt)
			if err != nil {
				return fmt.Errorf("%s: %w", msg.ErrorGenerateMessage, err)
			}

			// JIRA 이슈 번호 추가
			originalMessage := commitMessage
			commitMessage, _ = git.FormatCommitMessage(commitMessage, cfg.JiraIntegration)
			
			// JIRA 이슈가 추가되었으면 사용자에게 알림
			if originalMessage != commitMessage {
				branch, _ := git.GetCurrentBranch()
				jiraIssue := git.ExtractJiraIssue(branch)
				if jiraIssue != "" {
					color.Cyan(msg.JiraIssueDetected, jiraIssue)
					color.Green(msg.JiraIssueAdded)
				}
			}

			fmt.Println()
			color.Green(msg.NewMessageGenerated)
			fmt.Println()
			color.Yellow("%s", commitMessage)
			fmt.Println()
			continue

		case msg.PromptCancel:
			color.Yellow(msg.CommitCancelled)
			return nil
		default:
			// 예상치 못한 선택 (발생하지 않아야 함)
			return fmt.Errorf(msg.ErrorUnknownSelection, result)
		}
	}
}
