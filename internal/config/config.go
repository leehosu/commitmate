package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config는 애플리케이션 설정을 담는 구조체입니다
type Config struct {
	Provider string       `mapstructure:"provider"`
	Language string       `mapstructure:"language"`
	Template string       `mapstructure:"template"`
	OpenAI   OpenAIConfig `mapstructure:"openai"`
	Claude   ClaudeConfig `mapstructure:"claude"`
}

// OpenAIConfig는 OpenAI 관련 설정입니다
type OpenAIConfig struct {
	APIKey    string `mapstructure:"api_key"`
	Model     string `mapstructure:"model"`
	MaxTokens int    `mapstructure:"max_tokens"`
}

// ClaudeConfig는 Claude 관련 설정입니다
type ClaudeConfig struct {
	APIKey    string `mapstructure:"api_key"`
	Model     string `mapstructure:"model"`
	MaxTokens int    `mapstructure:"max_tokens"`
}

// Default는 기본 설정을 반환합니다
func Default() *Config {
	return &Config{
		Provider: "openai",
		Language: "en",
		Template: "conventional",
		OpenAI: OpenAIConfig{
			APIKey:    "",
			Model:     "gpt-4o",
			MaxTokens: 100, // Rate limit 방지를 위해 줄임
		},
		Claude: ClaudeConfig{
			APIKey:    "",
			Model:     "claude-3-5-sonnet-20241022",
			MaxTokens: 100, // Rate limit 방지를 위해 줄임
		},
	}
}

// Load는 설정 파일을 로드합니다
func Load() (*Config, error) {
	cfg := Default()

	// 홈 디렉토리 찾기
	home, err := os.UserHomeDir()
	if err != nil {
		return cfg, nil // 홈 디렉토리 없으면 기본값 사용
	}

	// Viper 설정
	configDir := filepath.Join(home, ".commitgen")
	configFile := filepath.Join(configDir, "config.yaml")

	// 설정 파일이 존재하는지 확인
	if _, err := os.Stat(configFile); err == nil {
		// 설정 파일 읽기
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err == nil {
			// 파일에서 설정 로드 (기본값 위에 덮어쓰기)
			if err := viper.Unmarshal(cfg); err != nil {
				return nil, fmt.Errorf("설정 파일 파싱 실패: %w", err)
			}
		}
	}

	// 환경변수로 덮어쓰기 (최우선)
	if apiKey := os.Getenv("COMMITGEN_OPENAI_API_KEY"); apiKey != "" {
		cfg.OpenAI.APIKey = apiKey
	}
	if apiKey := os.Getenv("COMMITGEN_CLAUDE_API_KEY"); apiKey != "" {
		cfg.Claude.APIKey = apiKey
	}
	if provider := os.Getenv("COMMITGEN_PROVIDER"); provider != "" {
		cfg.Provider = provider
	}

	return cfg, nil
}

// Save는 설정을 파일에 저장합니다
func Save(cfg *Config) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("홈 디렉토리를 찾을 수 없습니다: %w", err)
	}

	configDir := filepath.Join(home, ".commitgen")
	configFile := filepath.Join(configDir, "config.yaml")

	// 디렉토리가 없으면 생성
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("설정 디렉토리 생성 실패: %w", err)
	}

	viper.Set("provider", cfg.Provider)
	viper.Set("language", cfg.Language)
	viper.Set("template", cfg.Template)
	viper.Set("openai", cfg.OpenAI)
	viper.Set("claude", cfg.Claude)

	if err := viper.WriteConfigAs(configFile); err != nil {
		return fmt.Errorf("설정 파일 저장 실패: %w", err)
	}

	return nil
}

// GetConfigPath는 설정 파일 경로를 반환합니다
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".commitgen", "config.yaml"), nil
}
