package ai

import (
	"fmt"

	"github.com/leehosu/commitgen/internal/config"
)

// Client는 AI 클라이언트 인터페이스입니다
type Client interface {
	GenerateCommitMessage(systemPrompt, userPrompt string) (string, error)
}

// NewClient는 설정에 따라 적절한 AI 클라이언트를 생성합니다
func NewClient(cfg *config.Config) (Client, error) {
	switch cfg.Provider {
	case "openai":
		if cfg.OpenAI.APIKey == "" {
			return nil, fmt.Errorf("OpenAI API 키가 설정되지 않았습니다")
		}
		return NewOpenAIClient(cfg.OpenAI.APIKey, cfg.OpenAI.Model, cfg.OpenAI.MaxTokens), nil
	
	case "claude":
		if cfg.Claude.APIKey == "" {
			return nil, fmt.Errorf("Claude API 키가 설정되지 않았습니다")
		}
		return NewClaudeClient(cfg.Claude.APIKey, cfg.Claude.Model, cfg.Claude.MaxTokens), nil
	
	default:
		return nil, fmt.Errorf("지원하지 않는 제공자입니다: %s", cfg.Provider)
	}
}
