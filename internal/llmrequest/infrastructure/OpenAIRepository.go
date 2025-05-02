package infrastructure

import (
	"PetAi/internal/llmrequest"
	"PetAi/pkg/database"
)

type OpenAIRepository struct {
	db *database.DbConn
}

// Create a LLM request instance repository
func NewOpenAIRepository(dbcon *database.DbConn) llmrequest.LLMRepository {
	return &OpenAIRepository{db: dbcon}
}

// Send a new request in the LLM
func (repo *OpenAIRepository) SendRequest(p *llmrequest.Promt) (string, error) {
	// Get the id inserted in the database
	return "id", nil
}
