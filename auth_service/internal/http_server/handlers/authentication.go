package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/entities"
)

func (s RestServer) LogIn(w http.ResponseWriter, r *http.Request) error {
	input := &entities.LoginInput{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		return fmt.Errorf("decoding JSON error: %w", err)
	}

	result, err := s.service.SignIn(r.Context(), input)
	if err != nil {
		return fmt.Errorf("sign in error: %w", err)
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return fmt.Errorf("encoding JSON error: %w", err)
	}

	return nil
}

func (s RestServer) RefreshToken(w http.ResponseWriter, r *http.Request) error {
	input := &entities.RefreshTokensInput{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		return fmt.Errorf("decoding JSON error: %w", err)
	}

	result, err := s.service.RefreshTokens(r.Context(), input)
	if err != nil {
		return fmt.Errorf("sign in error: %w", err)
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return fmt.Errorf("encoding JSON error: %w", err)
	}

	return nil
}
