package service

import (
	"api-service-test/internal/models"
	"api-service-test/internal/repository"
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateQuestion(ctx context.Context, text string) (*models.Question, error) {
	return &models.Question{}, nil
}

func (s *Service) GetQuestions(ctx context.Context) ([]models.Question, error) {
	return []models.Question{}, nil
}

func (s *Service) GetQuestion(ctx context.Context, id uint) (*models.Question, error) {
	return &models.Question{}, nil
}

func (s *Service) DeleteQuestion(ctx context.Context, id uint) error {
	return nil
}

func (s *Service) CreateAnswer(ctx context.Context, questionID uint, userID uuid.UUID, text string) (*models.Answer, error) {
	return &models.Answer{}, nil
}

func (s *Service) GetAnswer(ctx context.Context, id uint) (*models.Answer, error) {
	return &models.Answer{}, nil
}

func (s *Service) DeleteAnswer(ctx context.Context, id uint) error {
	return nil
}
