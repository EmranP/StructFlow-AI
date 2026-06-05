package gemini

import (
	"context"

	"google.golang.org/genai"
)

type provider struct {
	client *genai.Client
	model  string
}

func New(
	ctx context.Context,
	apiKey string,
) (*provider, error) {

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

	return &provider{
		client: client,
		model:  "gemini-2.5-flash",
	}, nil
}

func (p *provider) GenerateStructure(
	ctx context.Context,
	prompt string,
) (string, error) {

	resp, err := p.client.Models.GenerateContent(
		ctx,
		p.model,
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
