package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/leehosu/commitmate/internal/config"
	"github.com/leehosu/commitmate/internal/i18n"
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
	Example: `  commitmate config set-key openai sk-xxxxx
  commitmate config set-key claude sk-ant-xxxxx`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]
		apiKey := args[1]

		cfg, err := config.Load()
		if err != nil {
			cfg = config.Default()
		}

		msg := i18n.GetMessages(cfg.UILanguage)

		if provider != "openai" && provider != "claude" {
			color.Red("❌ " + msg.ErrorInvalidProvider)
			os.Exit(1)
		}

		switch provider {
		case "openai":
			cfg.OpenAI.APIKey = apiKey
		case "claude":
			cfg.Claude.APIKey = apiKey
		}

		if err := config.Save(cfg); err != nil {
			color.Red("❌ "+msg.ErrorSaveConfig, err)
			os.Exit(1)
		}

		color.Green(msg.APIKeySaved, provider)
	},
}

var setProviderCmd = &cobra.Command{
	Use:   "set-provider [provider]",
	Short: "기본 AI 제공자를 설정합니다",
	Long:  `기본으로 사용할 AI 제공자(openai 또는 claude)를 설정합니다.`,
	Example: `  commitmate config set-provider openai
  commitmate config set-provider claude`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]

		cfg, err := config.Load()
		if err != nil {
			cfg = config.Default()
		}

		msg := i18n.GetMessages(cfg.UILanguage)

		if provider != "openai" && provider != "claude" {
			color.Red("❌ " + msg.ErrorInvalidProvider)
			os.Exit(1)
		}

		cfg.Provider = provider

		if err := config.Save(cfg); err != nil {
			color.Red("❌ "+msg.ErrorSaveConfig, err)
			os.Exit(1)
		}

		color.Green(msg.ProviderSet, provider)
	},
}

var setModelCmd = &cobra.Command{
	Use:   "set-model [provider] [model]",
	Short: "AI 모델을 설정합니다",
	Long:  `특정 제공자의 AI 모델을 변경합니다.`,
	Example: `  commitmate config set-model openai gpt-4o-mini
  commitmate config set-model claude claude-3-5-haiku-20241022`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]
		model := args[1]

		cfg, err := config.Load()
		if err != nil {
			cfg = config.Default()
		}

		msg := i18n.GetMessages(cfg.UILanguage)

		if provider != "openai" && provider != "claude" {
			color.Red("❌ " + msg.ErrorInvalidProvider)
			os.Exit(1)
		}

		switch provider {
		case "openai":
			cfg.OpenAI.Model = model
		case "claude":
			cfg.Claude.Model = model
		}

		if err := config.Save(cfg); err != nil {
			color.Red("❌ "+msg.ErrorSaveConfig, err)
			os.Exit(1)
		}

		color.Green(msg.ModelSet, provider, model)
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "현재 설정을 확인합니다",
	Long:  `현재 저장된 모든 설정을 출력합니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			// 설정 로드 실패 시 기본 언어로 메시지 표시
			defaultCfg := config.Default()
			msg := i18n.GetMessages(defaultCfg.UILanguage)
			color.Red("❌ "+msg.ErrorGetConfigPath, err)
			os.Exit(1)
		}

		msg := i18n.GetMessages(cfg.UILanguage)

		color.Cyan(msg.ConfigTitle)
		fmt.Println()

		color.White(msg.ConfigProvider, cfg.Provider)
		color.White(msg.ConfigCommitLanguage, cfg.CommitLanguage)
		color.White(msg.ConfigUILanguage, cfg.UILanguage)
		color.White(msg.ConfigTemplate, cfg.Template)
		fmt.Println()

		color.Yellow(msg.OpenAISettings)
		if cfg.OpenAI.APIKey != "" {
			maskedKey := maskAPIKey(cfg.OpenAI.APIKey)
			color.White(msg.APIKeyLabel, maskedKey)
		} else {
			color.White(msg.APIKeyNotSet)
		}
		color.White(msg.ModelLabel, cfg.OpenAI.Model)
		color.White(msg.MaxTokensLabel, cfg.OpenAI.MaxTokens)
		fmt.Println()

		color.Yellow(msg.ClaudeSettings)
		if cfg.Claude.APIKey != "" {
			maskedKey := maskAPIKey(cfg.Claude.APIKey)
			color.White(msg.APIKeyLabel, maskedKey)
		} else {
			color.White(msg.APIKeyNotSet)
		}
		color.White(msg.ModelLabel, cfg.Claude.Model)
		color.White(msg.MaxTokensLabel, cfg.Claude.MaxTokens)
	},
}

var setCommitLanguageCmd = &cobra.Command{
	Use:   "set-commit-language [language]",
	Short: "커밋 메시지 언어를 설정합니다",
	Long:  `AI가 생성하는 커밋 메시지의 언어를 설정합니다 (ko 또는 en).`,
	Example: `  commitmate config set-commit-language en
  commitmate config set-commit-language ko`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lang := args[0]

		cfg, err := config.Load()
		if err != nil {
			cfg = config.Default()
		}

		msg := i18n.GetMessages(cfg.UILanguage)

		if lang != "ko" && lang != "en" {
			color.Red("❌ " + msg.ErrorInvalidLanguage)
			os.Exit(1)
		}

		cfg.CommitLanguage = lang

		if err := config.Save(cfg); err != nil {
			color.Red("❌ "+msg.ErrorSaveConfig, err)
			os.Exit(1)
		}

		color.Green(msg.CommitLanguageSet, lang)
	},
}

var setUILanguageCmd = &cobra.Command{
	Use:   "set-ui-language [language]",
	Short: "UI 언어를 설정합니다",
	Long:  `CLI 인터페이스 메시지의 언어를 설정합니다 (ko 또는 en).`,
	Example: `  commitmate config set-ui-language en
  commitmate config set-ui-language ko`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lang := args[0]

		cfg, err := config.Load()
		if err != nil {
			cfg = config.Default()
		}

		// 현재 UI 언어로 메시지 가져오기
		msg := i18n.GetMessages(cfg.UILanguage)

		if lang != "ko" && lang != "en" {
			color.Red("❌ " + msg.ErrorInvalidLanguage)
			os.Exit(1)
		}

		cfg.UILanguage = lang

		if err := config.Save(cfg); err != nil {
			color.Red("❌ "+msg.ErrorSaveConfig, err)
			os.Exit(1)
		}

		// 변경된 언어로 성공 메시지 출력
		newMsg := i18n.GetMessages(lang)
		color.Green(newMsg.UILanguageSet, lang)
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
	configCmd.AddCommand(setCommitLanguageCmd)
	configCmd.AddCommand(setUILanguageCmd)
	configCmd.AddCommand(showCmd)
}
