package openai

import (
	"context"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type OpenAIProvider struct {
	client *openai.Client
}

func New(
	apiKey string,
) *OpenAIProvider {
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	return &OpenAIProvider{client: &client}
}

func (o *OpenAIProvider) Name() string {
	return "gpt"
}

func (o *OpenAIProvider) GenerateStructure(
	ctx context.Context,
	prompt string,
) (string, error) {

	resp, err := o.client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Model: "gpt-4o",
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(prompt),
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func (p *OpenAIProvider) IsAvailable(
	ctx context.Context,
) bool {
	_, err := p.GenerateStructure(
		ctx,
		"ping",
	)

	return err == nil
}
