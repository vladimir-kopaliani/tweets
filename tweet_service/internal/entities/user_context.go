package entities

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const (
	// UserContextKey is key for storing in `context.Context`
	UserContextKey ContextKey = "user_context"
)

type UserContext struct {
	AcessToken string
	UserID     string

	jwt.RegisteredClaims
}
