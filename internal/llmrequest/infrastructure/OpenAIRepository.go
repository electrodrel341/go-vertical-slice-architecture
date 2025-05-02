package infrastructure

import (
	"PetAi/internal/llmrequest"
	"PetAi/pkg/apperror"
	"PetAi/pkg/database"
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

const (
	ErrorNotFoundOpenAiApiKey string = "ERROR_NOT_FOUND_OPENAI_API_KEY"
)

type OpenAIRepository struct {
	db *database.DbConn

	apiKey string
	client *openai.Client
}

// Create a LLM request instance repository
func NewOpenAIRepository(dbcon *database.DbConn) (llmrequest.LLMRepository, error) {

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		err := apperror.ConfigNotFound(ErrorNotFoundOpenAiApiKey)
		log.Println("Ошибка инициализации OpenAI:", err)
		return nil, err
	}

	return &OpenAIRepository{
		db:     dbcon,
		apiKey: apiKey,
		client: openai.NewClient(apiKey),
	}, nil

	//return &OpenAIRepository{db: dbcon}
}

// Send a new request in the LLM
func (repo *OpenAIRepository) SendRequest(p *llmrequest.Promt) (string, error) {
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
		Model:    string(p.Model),
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
