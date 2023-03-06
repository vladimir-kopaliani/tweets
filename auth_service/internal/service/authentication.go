package service

import (
	"context"
	"fmt"
	"time"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/entities"

	apperrors "github.com/vladimir-kopaliani/tweets/auth_service/internal/app_errors"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtAccessTTL  = 5 * time.Minute
	jwtRefreshTTL = 3 * 24 * time.Hour
)

var method = jwt.SigningMethodHS512

func (s Service) SignIn(ctx context.Context, input *entities.LoginInput) (*entities.SignInResult, error) {
	userID, err := s.userService.CheckRegisteredUser(ctx, input.Login, input.Password)
	if err != nil {
		return nil, fmt.Errorf("checking registered user error: %w", err)
	}

	signedAccessToken, signedRefreshToken, err := s.generateTokens(userID)
	if err != nil {
		return nil, fmt.Errorf("tokens generation error: %w", err)
	}

	return &entities.SignInResult{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
		// User:         &user.User,
	}, nil
}

func (s Service) RefreshTokens(ctx context.Context, input *entities.RefreshTokensInput) (*entities.RefreshTokensResult, error) {
	refreshToken, err := jwt.Parse(input.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method of refresh token: %v", token.Header["alg"])
		}

		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("refresh token parsing error: %w", err)
	}

	if !refreshToken.Valid {
		return nil, apperrors.ErrJWTInvalid
	}

	claim, ok := refreshToken.Claims.(entities.UserContext)
	if !ok {
		return nil, apperrors.ErrJWTInvalid
	}

	signedAccessToken, signedRefreshToken, err := s.generateTokens(claim.UserID)
	if err != nil {
		return nil, fmt.Errorf("tokens generation error: %w", err)
	}

	return &entities.RefreshTokensResult{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}

func (s Service) generateTokens(userID string) (string, string, error) {
	now := time.Now()

	// generate access token with claim
	accessToken := jwt.NewWithClaims(
		method,
		entities.UserContext{
			UserID: userID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(now.Add(jwtAccessTTL)),
			},
		},
	)

	signedAccessToken, err := accessToken.SignedString(s.jwtSecret)
	if err != nil {
		return "", "", err
	}

	// generate refresh token with claim
	refreshToken := jwt.NewWithClaims(
		method,
		entities.UserContext{
			UserID: userID,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(now.Add(jwtRefreshTTL)),
			},
		},
	)

	signedRefreshToken, err := refreshToken.SignedString(s.jwtSecret)
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}
