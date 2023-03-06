package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (s RestServer) GetUserByID(w http.ResponseWriter, r *http.Request) error {
	userID := strings.TrimPrefix(r.URL.Path, "/")

	result, err := s.service.GetUserByID(r.Context(), userID)
	if err != nil {
		return fmt.Errorf("sign in error: %w", err)
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return fmt.Errorf("encoding JSON error: %w", err)
	}

	return nil
}
