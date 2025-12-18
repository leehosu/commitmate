package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/leehosu/commitgen/internal/ai"
	"github.com/leehosu/commitgen/internal/config"
	"github.com/leehosu/commitgen/internal/git"
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
	// 1. Git 저장소 확인
	color.Cyan("🔍 Git 저장소 확인 중...")
	if err := git.CheckRepository(); err != nil {
		return fmt.Errorf("Git 저장소가 아닙니다: %w", err)
	}

	// 2. Staged 변경사항 가져오기
	color.Cyan("📝 Staged 변경사항 분석 중...")
	diff, err := git.GetStagedDiff()
	if err != nil {
		return fmt.Errorf("Staged 변경사항을 가져올 수 없습니다: %w", err)
	}

	if diff == "" {
		return fmt.Errorf("커밋할 변경사항이 없습니다. 'git add' 명령어를 먼저 실행하세요")
	}

	// 3. 설정 로드
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("설정을 로드할 수 없습니다: %w", err)
	}

	// 명령줄에서 지정한 provider가 있으면 덮어쓰기
	if provider != "" {
		cfg.Provider = provider
	}

	// 4. AI 클라이언트 생성
	client, err := ai.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("AI 클라이언트를 생성할 수 없습니다: %w\n\n힌트: 'commitgen config set-key %s <API_KEY>' 명령어로 API 키를 설정하세요", cfg.Provider)
	}

	// 5. 커밋 메시지 생성
	color.Cyan("✨ AI가 커밋 메시지를 생성하는 중...")
	
	// 프롬프트 생성
	systemPrompt := prompt.GetSystemPrompt(cfg.Language, cfg.Template)
	userPrompt := prompt.GetUserPrompt(diff)
	
	commitMessage, err := client.GenerateCommitMessage(systemPrompt, userPrompt)
	if err != nil {
		return fmt.Errorf("커밋 메시지를 생성할 수 없습니다: %w", err)
	}

	// 6. 생성된 메시지 출력
	fmt.Println()
	color.Green("✨ 생성된 커밋 메시지:")
	fmt.Println()
	color.Yellow("%s", commitMessage)
	fmt.Println()

	// dry-run 모드면 여기서 종료
	if dryRun {
		color.Cyan("ℹ️  --dry-run 모드: 커밋하지 않습니다")
		return nil
	}

	// 7. 사용자 확인
	for {
		promptSelect := promptui.Select{
			Label: "이 커밋 메시지를 사용하시겠습니까?",
			Items: []string{
				"✓ Yes - 커밋 실행",
				"✎ Edit - 수정 후 커밋",
				"↻ Regenerate - 다시 생성",
				"✗ Cancel - 취소",
			},
		}

		_, result, err := promptSelect.Run()
		if err != nil {
			return fmt.Errorf("선택 중 오류 발생: %w", err)
		}

		switch result {
		case "✓ Yes - 커밋 실행":
			// 8. 커밋 실행
			color.Cyan("🚀 커밋 실행 중...")
			if err := git.Commit(commitMessage, noVerify); err != nil {
				return fmt.Errorf("커밋 실패: %w", err)
			}
			color.Green("✓ 커밋이 완료되었습니다!")
			return nil

		case "✎ Edit - 수정 후 커밋":
			// 수정 프롬프트
			promptEdit := promptui.Prompt{
				Label:   "커밋 메시지",
				Default: commitMessage,
			}
			editedMessage, err := promptEdit.Run()
			if err != nil {
				return fmt.Errorf("입력 중 오류 발생: %w", err)
			}
			
			// 수정된 메시지로 커밋
			color.Cyan("🚀 커밋 실행 중...")
			if err := git.Commit(editedMessage, noVerify); err != nil {
				return fmt.Errorf("커밋 실패: %w", err)
			}
			color.Green("✓ 커밋이 완료되었습니다!")
			return nil

		case "↻ Regenerate - 다시 생성":
			// 다시 생성
			color.Cyan("✨ 커밋 메시지를 다시 생성하는 중...")
			commitMessage, err = client.GenerateCommitMessage(systemPrompt, userPrompt)
			if err != nil {
				return fmt.Errorf("커밋 메시지를 생성할 수 없습니다: %w", err)
			}
			
			fmt.Println()
			color.Green("✨ 새로 생성된 커밋 메시지:")
			fmt.Println()
			color.Yellow("%s", commitMessage)
			fmt.Println()
			continue

		case "✗ Cancel - 취소":
			color.Yellow("커밋이 취소되었습니다")
			return nil
		}
	}
}
