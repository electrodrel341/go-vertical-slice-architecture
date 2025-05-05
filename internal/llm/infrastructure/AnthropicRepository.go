package infrastructure

import (
	"PetAi/internal/llm"
	"PetAi/pkg/database"
)

type AnthropicRepository struct {
	db *database.DbConn
}

// Create a Anthropic request instance repository
func NewAnthropicRepository(dbcon *database.DbConn) llm.LLMRepository {
	return &AnthropicRepository{db: dbcon}
}

// Send a new request in the LLM
func (repo *AnthropicRepository) SendRequest(p *llm.Promt) (string, error) {
	// Get the id inserted in the database
	return "id2", nil
}
