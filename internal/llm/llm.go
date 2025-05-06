package llm

import (
	"PetAi/pkg/apperror"
	"strings"
)

type AIApiProvider int
type AIModel int

// Enumeration of product categories
const (
	UndefinedProvider AIApiProvider = iota
	OpenAi
	Yandex
	Anthropic
	GoogleCloud
)

// Срез всех значений
var AllProviders = []AIApiProvider{
	OpenAi,
	Anthropic,
	GoogleCloud,
}

func (p AIApiProvider) String() string {
	switch p {
	case OpenAi:
		return "OpenAi"
	case Yandex:
		return "Yandex"
	case Anthropic:
		return "Anthropic"
	case GoogleCloud:
		return "GoogleCloud"
	default:
		return "UndefinedProvider"
	}
}

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

const (
	UndefinedModel AIModel = iota
	O1
	O3
	O4
	O4mini
	Claude37sonnet
	Googledefaultmodel
	Yandexdefaultmodel
)

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

var ProviderToModels = map[AIApiProvider][]AIModel{
	OpenAi:      {O1, O3, O4, O4mini},
	Yandex:      {Yandexdefaultmodel},
	Anthropic:   {Claude37sonnet},
	GoogleCloud: {Googledefaultmodel},
}

var ProviderDefaultModel = map[AIApiProvider]AIModel{
	OpenAi:      O4mini,
	Yandex:      Yandexdefaultmodel,
	Anthropic:   Claude37sonnet,
	GoogleCloud: Googledefaultmodel,
}

func IsModelSupported(provider AIApiProvider, model AIModel) bool {
	models, ok := ProviderToModels[provider]
	if !ok {
		return false
	}
	for _, m := range models {
		if m == model {
			return true
		}
	}
	return false
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

func NewPromt(
	message string,
	model *AIModel,
	maxTokens *int,
	temperature *float32,
	topP *float32,
	systemMessage string,
	messagesHistory []string,
	providerParams map[string]interface{},
) Promt {
	var selectedModel AIModel
	if model != nil {
		selectedModel = *model
	}

	return Promt{
		RequestMessage:         message,
		Model:                  selectedModel,
		Max_tokens:             maxTokens,
		Temperature:            temperature,
		TopP:                   topP,
		SystemMessage:          systemMessage,
		MessagesHistory:        messagesHistory,
		ProviderSpecificParams: providerParams,
	}
}

func NewLLMRequest(provider AIApiProvider, promt Promt) *LLMRequest {
	model := promt.Model

	if model == UndefinedModel {
		// Модель не задана — берём по умолчанию
		defaultModel, ok := ProviderDefaultModel[provider]
		if !ok {
			return nil
		}
		promt.Model = defaultModel
	} else if !IsModelSupported(provider, model) {
		return nil
	}

	return &LLMRequest{
		AIApiProvider: provider,
		Promt:         promt,
	}
}
