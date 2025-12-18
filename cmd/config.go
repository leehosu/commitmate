package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/leehosu/commitgen/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "설정을 관리합니다",
	Long:  `API 키, 제공자, 모델 등의 설정을 관리합니다.`,
}

var setKeyCmd = &cobra.Command{
	Use:   "set-key [provider] [api-key]",
	Short: "API 키를 설정합니다",
	Long:  `OpenAI 또는 Claude의 API 키를 설정합니다.`,
	Example: `  commitgen config set-key openai sk-xxxxx
  commitgen config set-key claude sk-ant-xxxxx`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]
		apiKey := args[1]

		if provider != "openai" && provider != "claude" {
			color.Red("❌ 잘못된 제공자입니다. 'openai' 또는 'claude'를 사용하세요")
			os.Exit(1)
		}

		cfg, err := config.Load()
		if err != nil {
			cfg = config.Default()
		}

		switch provider {
		case "openai":
			cfg.OpenAI.APIKey = apiKey
		case "claude":
			cfg.Claude.APIKey = apiKey
		}

		if err := config.Save(cfg); err != nil {
			color.Red("❌ 설정 저장 실패: %v", err)
			os.Exit(1)
		}

		color.Green("✓ %s API 키가 저장되었습니다", provider)
	},
}

var setProviderCmd = &cobra.Command{
	Use:   "set-provider [provider]",
	Short: "기본 AI 제공자를 설정합니다",
	Long:  `기본으로 사용할 AI 제공자(openai 또는 claude)를 설정합니다.`,
	Example: `  commitgen config set-provider openai
  commitgen config set-provider claude`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]

		if provider != "openai" && provider != "claude" {
			color.Red("❌ 잘못된 제공자입니다. 'openai' 또는 'claude'를 사용하세요")
			os.Exit(1)
		}

		cfg, err := config.Load()
		if err != nil {
			cfg = config.Default()
		}

		cfg.Provider = provider

		if err := config.Save(cfg); err != nil {
			color.Red("❌ 설정 저장 실패: %v", err)
			os.Exit(1)
		}

		color.Green("✓ 기본 제공자가 %s로 설정되었습니다", provider)
	},
}

var setModelCmd = &cobra.Command{
	Use:   "set-model [provider] [model]",
	Short: "AI 모델을 설정합니다",
	Long:  `특정 제공자의 AI 모델을 변경합니다.`,
	Example: `  commitgen config set-model openai gpt-4o-mini
  commitgen config set-model claude claude-3-5-haiku-20241022`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]
		model := args[1]

		if provider != "openai" && provider != "claude" {
			color.Red("❌ 잘못된 제공자입니다. 'openai' 또는 'claude'를 사용하세요")
			os.Exit(1)
		}

		cfg, err := config.Load()
		if err != nil {
			cfg = config.Default()
		}

		switch provider {
		case "openai":
			cfg.OpenAI.Model = model
		case "claude":
			cfg.Claude.Model = model
		}

		if err := config.Save(cfg); err != nil {
			color.Red("❌ 설정 저장 실패: %v", err)
			os.Exit(1)
		}

		color.Green("✓ %s 모델이 %s로 설정되었습니다", provider, model)
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "현재 설정을 확인합니다",
	Long:  `현재 저장된 모든 설정을 출력합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			color.Red("❌ 설정을 불러올 수 없습니다: %v", err)
			os.Exit(1)
		}

		color.Cyan("📋 현재 설정:")
		fmt.Println()
		
		color.White("기본 제공자: %s", cfg.Provider)
		color.White("언어: %s", cfg.Language)
		color.White("템플릿: %s", cfg.Template)
		fmt.Println()

		color.Yellow("OpenAI 설정:")
		if cfg.OpenAI.APIKey != "" {
			maskedKey := maskAPIKey(cfg.OpenAI.APIKey)
			color.White("  API 키: %s", maskedKey)
		} else {
			color.White("  API 키: (설정되지 않음)")
		}
		color.White("  모델: %s", cfg.OpenAI.Model)
		color.White("  Max Tokens: %d", cfg.OpenAI.MaxTokens)
		fmt.Println()

		color.Yellow("Claude 설정:")
		if cfg.Claude.APIKey != "" {
			maskedKey := maskAPIKey(cfg.Claude.APIKey)
			color.White("  API 키: %s", maskedKey)
		} else {
			color.White("  API 키: (설정되지 않음)")
		}
		color.White("  모델: %s", cfg.Claude.Model)
		color.White("  Max Tokens: %d", cfg.Claude.MaxTokens)
	},
}

func maskAPIKey(key string) string {
	if len(key) <= 8 {
		return "***"
	}
	return key[:4] + "..." + key[len(key)-4:]
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setKeyCmd)
	configCmd.AddCommand(setProviderCmd)
	configCmd.AddCommand(setModelCmd)
	configCmd.AddCommand(showCmd)
}
