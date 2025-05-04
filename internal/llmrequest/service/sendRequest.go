package service

import (
	"PetAi/internal/llmrequest"
	"PetAi/pkg/apperror"
	"log"
)

type SendLLMRequestService struct {
	router llmrequest.LLMRepositoryRouter
}

// Create a new llm request service use case instance
/*func NewSendLLMRequestService(repository llmrequest.LLMRepository) *SendLLMRequestService {
	// return the pointer to product service
	return &SendLLMRequestService{
		llmRepository: repository,
	}
}*/

func NewSendLLMRequestService(router llmrequest.LLMRepositoryRouter) *SendLLMRequestService {
	return &SendLLMRequestService{
		router: router,
	}
}

// Create a new product and store the product into the database
func (service *SendLLMRequestService) SendRequest(p *llmrequest.LLMRequest) (string, error) {
	repo, err := service.router.GetRepository(p.AIApiProvider)

	response, err := repo.SendRequest(&p.Promt)

	if err != nil {
		log.Println(err)
		err := apperror.InternalServerError(err)
		return "", err
	}

	// product created successfuly
	return response, nil
}
