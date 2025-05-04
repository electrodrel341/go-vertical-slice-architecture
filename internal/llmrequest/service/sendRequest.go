package service

import (
	"PetAi/internal/llmrequest"
	"PetAi/pkg/apperror"
	"github.com/rs/zerolog"
)

type SendLLMRequestService struct {
	router llmrequest.LLMRepositoryRouter
	logger zerolog.Logger
}

// Create a new llm request service use case instance
/*func NewSendLLMRequestService(repository llmrequest.LLMRepository) *SendLLMRequestService {
	// return the pointer to product service
	return &SendLLMRequestService{
		llmRepository: repository,
	}
}*/

func NewSendLLMRequestService(
	router llmrequest.LLMRepositoryRouter,
	logger zerolog.Logger, // добавили логгер
) *SendLLMRequestService {
	return &SendLLMRequestService{
		router: router,
		logger: logger,
	}
}

// Create a new product and store the product into the database
func (service *SendLLMRequestService) SendRequest(p *llmrequest.LLMRequest) (string, error) {

	repo, err := service.router.GetRepository(p.AIApiProvider)
	response, err := repo.SendRequest(&p.Promt)

	if err != nil {
		service.logger.Error().Str("provider", string(p.AIApiProvider)).Msg("Error sending LLM request")
		err := apperror.InternalServerError(err)
		return "", err
	}

	return response, nil
}
