package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/leehosu/commitgen/internal/ai"
	"github.com/leehosu/commitgen/internal/config"
	"github.com/leehosu/commitgen/internal/git"
	"github.com/leehosu/commitgen/internal/i18n"
	"github.com/leehosu/commitgen/internal/prompt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "AI가 생성한 커밋 메시지로 커밋합니다",
	Long:  `Staged 변경사항을 분석하여 AI가 자동으로 커밋 메시지를 생성하고 커밋합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runCommit(); err != nil {
			color.Red("❌ 오류: %v", err)
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
		return fmt.Errorf("설정을 로드할 수 없습니다: %w", err)
	}

	// UI 메시지 가져오기
	msg := i18n.GetMessages(cfg.UILanguage)

	// 1. Git 저장소 확인
	color.Cyan(msg.CheckingRepository)
	if err := git.CheckRepository(); err != nil {
		return fmt.Errorf(msg.ErrorNotGitRepo+": %w", err)
	}

	// 2. Staged 변경사항 가져오기
	color.Cyan(msg.AnalyzingStagedChanges)
	diff, err := git.GetStagedDiff()
	if err != nil {
		return fmt.Errorf(msg.ErrorNoStagedChanges+": %w", err)
	}

	if diff == "" {
		return fmt.Errorf(msg.NoChanges)
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
		return fmt.Errorf(msg.ErrorCreateClient+"\n\n"+msg.HintSetAPIKey, err, cfg.Provider)
	}

	// 5. 커밋 메시지 생성
	color.Cyan(msg.GeneratingMessage)

	// 프롬프트 생성 (CommitLanguage 사용)
	systemPrompt := prompt.GetSystemPrompt(cfg.CommitLanguage, cfg.Template)
	userPrompt := prompt.GetUserPrompt(diff)

	commitMessage, err := client.GenerateCommitMessage(systemPrompt, userPrompt)
	if err != nil {
		return fmt.Errorf(msg.ErrorGenerateMessage+": %w", err)
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
			return fmt.Errorf(msg.ErrorSelectFailed+": %w", err)
		}

		switch result {
		case msg.PromptYes:
			// 8. 커밋 실행
			color.Cyan(msg.Committing)
			if err := git.Commit(commitMessage, noVerify); err != nil {
				return fmt.Errorf(msg.ErrorCommitFailed+": %w", err)
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
				return fmt.Errorf(msg.ErrorInputFailed+": %w", err)
			}

			// 수정된 메시지로 커밋
			color.Cyan(msg.Committing)
			if err := git.Commit(editedMessage, noVerify); err != nil {
				return fmt.Errorf(msg.ErrorCommitFailed+": %w", err)
			}
			color.Green(msg.CommitSuccess)
			return nil

		case msg.PromptRegenerate:
			// 다시 생성
			color.Cyan(msg.RegeneratingMessage)
			commitMessage, err = client.GenerateCommitMessage(systemPrompt, userPrompt)
			if err != nil {
				return fmt.Errorf(msg.ErrorGenerateMessage+": %w", err)
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
			return fmt.Errorf("알 수 없는 선택: %s", result)
		}
	}
}
