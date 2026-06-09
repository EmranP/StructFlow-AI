package claude

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type ClaudeProvider struct {
	client *anthropic.Client
}

func New(
	apiKey string,
) *ClaudeProvider {
	client := anthropic.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &ClaudeProvider{client: &client}
}

func (c *ClaudeProvider) Name() string {
	return "claude"
}

func (c *ClaudeProvider) GenerateStructure(
	ctx context.Context,
	prompt string,
) (string, error) {
	resp, err := c.client.Messages.New(
		ctx,
		anthropic.MessageNewParams{
			Model:     anthropic.ModelClaudeSonnet4_5,
			MaxTokens: 8000,
			Messages: []anthropic.MessageParam{
				anthropic.NewUserMessage(
					anthropic.NewTextBlock(prompt),
				),
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Content[0].Text, nil
}

func (p *ClaudeProvider) IsAvailable(
	ctx context.Context,
) bool {
	_, err := p.GenerateStructure(
		ctx,
		"ping",
	)

	return err == nil
}
