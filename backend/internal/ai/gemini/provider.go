package gemini

import (
	"context"

	"google.golang.org/genai"
)

type GeminiProvider struct {
	client *genai.Client
	model  string
}

func New(
	ctx context.Context,
	apiKey string,
) (*GeminiProvider, error) {

	client, err := genai.NewClient(
		ctx,
		&genai.ClientConfig{
			APIKey:  apiKey,
			Backend: genai.BackendGeminiAPI,
		},
	)

	if err != nil {
		return nil, err
	}

	return &GeminiProvider{
		client: client,
		model:  "gemini-2.5-flash",
	}, nil
}

func (g *GeminiProvider) Name() string {
	return "gemini"
}

func (p *GeminiProvider) GenerateStructure(
	ctx context.Context,
	prompt string,
) (string, error) {

	resp, err := p.client.Models.GenerateContent(
		ctx,
		p.model,
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			MaxOutputTokens: 8000,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}

func (p *GeminiProvider) IsAvailable(
	ctx context.Context,
) bool {
	_, err := p.GenerateStructure(
		ctx,
		"ping",
	)

	return err == nil
}
