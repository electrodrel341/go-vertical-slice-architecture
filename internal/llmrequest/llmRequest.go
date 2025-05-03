package llmrequest

import (
	"PetAi/pkg/apperror"
	"strings"
)

type AIApiProvider int
type AIModel int

// Enumeration of product categories
const (
	OpenAi AIApiProvider = iota
	Yandex
	Anthropic
	GoogleCloud
	UndefinedProvider
)

const (
	O1 AIModel = iota
	O3
	O4
	O4mini
	Claude37sonnet
	Googledefaultmodel
	Yandexdefaultmodel
	UndefinedModel
)

/*// String representation of the ProductCategory
func (p AIApiProvider) String() string {
	return [...]string{
		"openai",
		"yandex",
		"anthropic-ai",
		"google-cloud",
	}[p]
}*/

// Parse ProductCategory converts a string to ProductCategory
func ParseAIApiProvider(s string) (AIApiProvider, error) {
	switch strings.ToLower(s) {
	case "openai":
		return OpenAi, nil
	case "yandex":
		return Yandex, nil
	case "anthropic-ai":
		return Anthropic, nil
	case "google-cloud":
		return GoogleCloud, nil
	default:
		{
			err := apperror.EntityNotFound(apperror.ErrorWrongAIApiProvider)
			return UndefinedProvider, err
		}

	}
}

func (p AIModel) String() string {
	return [...]string{
		"o1",
		"o3",
		"o4",
		"gpt-4o-mini",
		"yandex",
		"claude-3.7-sonnet",
		"google-cloud",
	}[p]
}

func ParseAIModel(s string) (AIModel, error) {
	switch strings.ToLower(s) {
	case "o1":
		return O1, nil
	case "o3":
		return O3, nil
	case "o4":
		return O4, nil
	case "gpt-4o-mini":
		return O4mini, nil
	case "yandex":
		return Yandexdefaultmodel, nil
	case "claude-3.7-sonnet":
		return Claude37sonnet, nil
	case "google-cloud":
		return Googledefaultmodel, nil
	default:
		{
			err := apperror.EntityNotFound(apperror.ErrorWrongAIModel)
			return UndefinedModel, err
		}

	}
}

type LLMRequest struct {
	AIApiProvider AIApiProvider
	Promt         Promt
}

// Promt Domain
type Promt struct {
	RequestMessage         string
	Model                  AIModel
	Max_tokens             *int
	Temperature            *float32
	TopP                   *float32
	SystemMessage          string
	MessagesHistory        []string
	ProviderSpecificParams map[string]interface{}
}

// Create a new product instance
func New(provider AIApiProvider, message string, model AIModel) (*LLMRequest, error) {
	return &LLMRequest{
		AIApiProvider: provider,
		Promt: Promt{
			RequestMessage: message,
			Model:          model,
		},
	}, nil
}
