package service

import (
	"api-service-test/internal/models"
	"api-service-test/internal/repository"
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateQuestion(ctx context.Context, text string) (*models.Question, error) {
	return s.repo.CreateQuestion(ctx, text)
}

func (s *Service) GetQuestions(ctx context.Context) ([]models.Question, error) {
	return s.repo.GetQuestions(ctx)
}

func (s *Service) GetQuestion(ctx context.Context, id uint) (*models.Question, error) {
	question, err := s.repo.GetQuestionByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, errors.New("not found")
	}
	return question, nil
}

func (s *Service) DeleteQuestion(ctx context.Context, id uint) error {
	return s.repo.DeleteQuestion(ctx, id)
}

func (s *Service) CreateAnswer(ctx context.Context, questionID uint, userID uuid.UUID, text string) (*models.Answer, error) {
	return s.repo.CreateAnswer(ctx, questionID, userID, text)
}

func (s *Service) GetAnswer(ctx context.Context, id uint) (*models.Answer, error) {
	answer, err := s.repo.GetAnswerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if answer == nil {
		return nil, errors.New("not found")
	}
	return answer, nil
}

func (s *Service) DeleteAnswer(ctx context.Context, id uint) error {
	return s.repo.DeleteAnswer(ctx, id)
}
