package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// ClaudeClient는 Anthropic Claude API를 사용하는 클라이언트입니다
type ClaudeClient struct {
	apiKey    string
	model     string
	maxTokens int
	httpClient *http.Client
}

// NewClaudeClient는 새로운 Claude 클라이언트를 생성합니다
func NewClaudeClient(apiKey, model string, maxTokens int) *ClaudeClient {
	return &ClaudeClient{
		apiKey:     apiKey,
		model:      model,
		maxTokens:  maxTokens,
		httpClient: &http.Client{},
	}
}

type claudeRequest struct {
	Model       string          `json:"model"`
	MaxTokens   int             `json:"max_tokens"`
	Messages    []claudeMessage `json:"messages"`
	System      string          `json:"system,omitempty"`
	Temperature float64         `json:"temperature,omitempty"`
}

type claudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type claudeResponse struct {
	Content []claudeContent `json:"content"`
	Error   *claudeError    `json:"error,omitempty"`
}

type claudeContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type claudeError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// GenerateCommitMessage는 Claude API를 사용하여 커밋 메시지를 생성합니다
func (c *ClaudeClient) GenerateCommitMessage(systemPrompt, userPrompt string) (string, error) {
	reqBody := claudeRequest{
		Model:     c.model,
		MaxTokens: c.maxTokens,
		Messages: []claudeMessage{
			{
				Role:    "user",
				Content: userPrompt,
			},
		},
		System:      systemPrompt,
		Temperature: 0.7,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("JSON 직렬화 실패: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("HTTP 요청 생성 실패: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("응답 읽기 실패: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Claude API 오류 (상태 코드: %d): %s", resp.StatusCode, string(body))
	}

	var claudeResp claudeResponse
	if err := json.Unmarshal(body, &claudeResp); err != nil {
		return "", fmt.Errorf("JSON 파싱 실패: %w", err)
	}

	if claudeResp.Error != nil {
		return "", fmt.Errorf("Claude API 오류: %s", claudeResp.Error.Message)
	}

	if len(claudeResp.Content) == 0 {
		return "", fmt.Errorf("Claude API 응답이 비어있습니다")
	}

	var result strings.Builder
	for _, content := range claudeResp.Content {
		if content.Type == "text" {
			result.WriteString(content.Text)
		}
	}

	return strings.TrimSpace(result.String()), nil
}
