package ai

import "context"

type Manager struct {
	providers map[string]Provider
}

func NewManager(
	providers ...Provider,
) *Manager {

	m := &Manager{
		providers: make(
			map[string]Provider,
		),
	}

	for _, provider := range providers {
		m.providers[provider.Name()] = provider
	}

	return &Manager{
		providers: m.providers,
	}
}

func (m *Manager) Get(
	name string,
) (Provider, bool) {

	p, ok := m.providers[name]

	return p, ok
}

func (m *Manager) Models() []ModelResponse {

	models := make(
		[]ModelResponse,
		0,
	)

	for _, provider := range m.providers {

		available := true

		if isPing := provider.IsAvailable(
			context.Background(),
		); !isPing {

			available = false
		}

		models = append(
			models,
			ModelResponse{
				ID: provider.Name(),

				Name: provider.Name(),

				Available: available,
			},
		)
	}

	return models
}
