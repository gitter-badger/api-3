package db

import (
	"github.com/lavab/api/models"
)

// TokensTable implements the CRUD interface for tokens
type TokensTable struct {
	RethinkCRUD
}

// GetToken returns a token with specified name
func (s *TokensTable) GetToken(id string) (*models.Token, error) {
	var result models.Token

	if err := s.FindFetchOne(id, &result); err != nil {
		return nil, err
	}

	return &result, nil
}