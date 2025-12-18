package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	provider string
	dryRun   bool
	noVerify bool
)

// rootCmd는 인자 없이 호출될 때의 기본 명령어입니다
var rootCmd = &cobra.Command{
	Use:   "commitgen",
	Short: "AI 기반 Git 커밋 메시지 자동 생성 도구",
	Long: `commitgen은 AI를 활용하여 Git의 staged 변경사항을 분석하고
자동으로 의미있는 커밋 메시지를 생성하여 커밋하는 도구입니다.

OpenAI GPT와 Anthropic Claude를 지원하며,
Conventional Commits 형식을 따릅니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 기본 동작: commitCmd 실행
		commitCmd.Run(cmd, args)
	},
}

// Execute는 루트 명령어를 실행합니다
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// 전역 플래그
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "설정 파일 경로 (기본값: $HOME/.commitgen/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&provider, "provider", "p", "", "AI 제공자 (openai 또는 claude)")

	// 로컬 플래그
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "커밋 메시지만 생성하고 커밋하지 않음")
	rootCmd.Flags().BoolVar(&noVerify, "no-verify", false, "git commit hooks 무시")
}

// initConfig는 설정 파일과 환경변수를 읽습니다
func initConfig() {
	if cfgFile != "" {
		// 명령줄에서 지정한 설정 파일 사용
		viper.SetConfigFile(cfgFile)
	} else {
		// 홈 디렉토리 찾기
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// .commitgen 디렉토리에서 설정 파일 찾기
		configDir := home + "/.commitgen"
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	// 환경변수 읽기 (COMMITGEN_ 접두사)
	viper.SetEnvPrefix("COMMITGEN")
	viper.AutomaticEnv()

	// 설정 파일 읽기 (없어도 에러 무시)
	if err := viper.ReadInConfig(); err == nil {
		// 설정 파일 찾음 (디버그용)
	}
}
