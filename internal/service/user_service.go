package service

import (
	"context"
	"fmt"
	"github.com/joramuns/shop/internal/kafka"
	"github.com/joramuns/shop/internal/models"
	"github.com/joramuns/shop/internal/repository"
)

type UserService struct {
	repo     *repository.UserRepository
	producer *kafka.Producer
}

func NewUserService(repo *repository.UserRepository, producer *kafka.Producer) *UserService {
	return &UserService{repo: repo, producer: producer}
}

func (s *UserService) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("error creating user service: %v", err)
	}
	if err := s.producer.SendMessage(ctx, createdUser); err != nil {
		return nil, fmt.Errorf("error sending message service: %v", err)
	}
	return createdUser, nil
}
