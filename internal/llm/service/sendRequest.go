package service

import (
	"PetAi/internal/llm"
	"PetAi/pkg/apperror"
	"github.com/rs/zerolog"
)

type SendLLMRequestService struct {
	router llm.LLMRepositoryRouter
	logger zerolog.Logger
}

func NewSendLLMRequestService(
	router llm.LLMRepositoryRouter,
	logger zerolog.Logger,
) *SendLLMRequestService {
	return &SendLLMRequestService{
		router: router,
		logger: logger,
	}
}

func (service *SendLLMRequestService) SendRequest(p *llm.LLMRequest) (string, error) {

	repo, err := service.router.GetRepository(p.AIApiProvider)
	response, err := repo.SendRequest(&p.Promt)

	if err != nil {
		service.logger.Error().Str("provider", string(p.AIApiProvider)).Msg("Error sending LLM request")
		err := apperror.InternalServerError(err)
		return "", err
	}

	return response, nil
}

func (service *SendLLMRequestService) SendRequestByAllProviders(p *llm.Promt) error {

	/*	for _, provider := range llm.AllProviders {
		repo, err := service.router.GetRepository(provider)
		response, err := repo.SendRequest(&p.Promt)

		if err != nil {
			service.logger.Error().Str("provider", string(p.AIApiProvider)).Msg("Error sending LLM request")
			err := apperror.InternalServerError(err)
			return "", err
		}
		return response, nil
	}*/

	return nil
}
