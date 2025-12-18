package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// OpenAIClient는 OpenAI API를 사용하는 클라이언트입니다
type OpenAIClient struct {
	client    *openai.Client
	model     string
	maxTokens int
}

// NewOpenAIClient는 새로운 OpenAI 클라이언트를 생성합니다
func NewOpenAIClient(apiKey, model string, maxTokens int) *OpenAIClient {
	client := openai.NewClient(apiKey)
	return &OpenAIClient{
		client:    client,
		model:     model,
		maxTokens: maxTokens,
	}
}

// GenerateCommitMessage는 OpenAI API를 사용하여 커밋 메시지를 생성합니다
func (c *OpenAIClient) GenerateCommitMessage(systemPrompt, userPrompt string) (string, error) {
	ctx := context.Background()

	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: c.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userPrompt,
				},
			},
			MaxTokens:   c.maxTokens,
			Temperature: 0.7,
		},
	)

	if err != nil {
		// Rate limit 에러인지 확인
		if strings.Contains(err.Error(), "rate_limit") || strings.Contains(err.Error(), "quota") {
			return "", fmt.Errorf("OpenAI API rate limit 초과: 잠시 후 다시 시도하거나, 더 작은 변경사항으로 나눠서 커밋하세요\n\n원본 에러: %w", err)
		}
		return "", fmt.Errorf("OpenAI API 호출 실패: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("OpenAI API 응답이 비어있습니다")
	}

	message := strings.TrimSpace(resp.Choices[0].Message.Content)
	return message, nil
}
