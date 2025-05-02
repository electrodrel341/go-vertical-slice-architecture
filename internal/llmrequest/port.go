package llmrequest

type LLMRepository interface {
	SendRequest(p *Promt) (string, error)
}

type LLMRepositoryRouter interface {
	GetRepository(provider AIApiProvider) (LLMRepository, error)
}
