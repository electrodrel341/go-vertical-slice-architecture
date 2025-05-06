package infrastructure

import (
	"PetAi/internal/llm"
	"PetAi/pkg/apperror"
	appConfig "PetAi/pkg/config"
	"PetAi/pkg/database"
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"net"
	"net/http"
	"net/url"
	"time"
)

type OpenAIRepository struct {
	db *database.DbConn

	apiKey string
	client *openai.Client
}

// Create a LLM request instance repository
func NewOpenAIRepository(dbcon *database.DbConn) (llm.LLMRepository, error) {

	apiKey := appConfig.Get().LLMConfig.OpenAIAPIKey

	if apiKey == "" {
		err := apperror.ConfigNotFound(apperror.ErrorNotFoundOpenAiApiKey)
		return nil, err
	}

	config := openai.DefaultConfig(apiKey)

	proxyURL, _ := url.Parse(appConfig.Get().ProxyConfig.Url)

	// Создаем Transport с прокси
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	// Создаем клиент с кастомным Transport
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   60 * time.Second,
	}

	config.HTTPClient = httpClient
	client := openai.NewClientWithConfig(config)

	return &OpenAIRepository{
		db:     dbcon,
		apiKey: apiKey,
		client: client,
	}, nil

}

// Send a new request in the LLM
func (repo *OpenAIRepository) SendRequest(p *llm.Promt) (string, error) {
	// Формируем сообщения для чата
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: p.SystemMessage,
		},
	}

	// Добавляем историю сообщений, если она есть
	for i := 0; i < len(p.MessagesHistory); i += 2 {
		if i < len(p.MessagesHistory) {
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: p.MessagesHistory[i],
			})
		}
		if i+1 < len(p.MessagesHistory) {
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: p.MessagesHistory[i+1],
			})
		}
	}

	// Добавляем текущий запрос пользователя
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: p.RequestMessage,
	})

	// Формируем параметры запроса
	req := openai.ChatCompletionRequest{
		Model:    p.Model.String(),
		Messages: messages,
	}

	// Добавляем опциональные параметры, если они указаны
	if p.Max_tokens != nil {
		req.MaxTokens = *p.Max_tokens
	}
	if p.Temperature != nil {
		req.Temperature = *p.Temperature
	}
	if p.TopP != nil {
		req.TopP = *p.TopP
	}

	// Выполняем запрос к OpenAI
	resp, err := repo.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("ошибка при вызове OpenAI API: %w", err)
	}

	// Сохраняем результат в БД или выполняем другие операции
	// repo.db.SaveResponse(...)

	// Возвращаем ответ модели
	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("получен пустой ответ от модели")
}
