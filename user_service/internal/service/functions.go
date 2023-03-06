package service

import (
	"context"
	"fmt"
	"time"

	"github.com/vladimir-kopaliani/tweets/user_service/internal/entities"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s Service) RegisterUser(ctx context.Context, user *entities.FullUserInfo) (*entities.User, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("coudn't generate UUID: %w", err)
	}

	user.ID = uuid.String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("coudn't hash password: %w", err)
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	err = s.pgRepository.SaveNewUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("saving new user error: %w", err)
	}

	return &user.User, err
}

func (s Service) GetUserByID(ctx context.Context, id string) (*entities.User, error) {
	user, err := s.pgRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s Service) getFullUserInfoByEmail(ctx context.Context, email string) (*entities.FullUserInfo, error) {
	user, err := s.pgRepository.GetFullUserInfoByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
