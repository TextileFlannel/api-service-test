package repository

import (
	"api-service-test/internal/models"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateQuestion(ctx context.Context, text string) (*models.Question, error) {
	question := &models.Question{
		Text: text,
	}

	if err := r.db.WithContext(ctx).Create(question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func (r *Repository) GetQuestions(ctx context.Context) ([]models.Question, error) {
	var questions []models.Question

	if err := r.db.WithContext(ctx).Order("created_at desc").Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *Repository) GetQuestionByID(ctx context.Context, id uint) (*models.Question, error) {
	var question models.Question
	err := r.db.WithContext(ctx).Preload("Answers").First(&question, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &question, nil
}

func (r *Repository) DeleteQuestion(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Question{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Question not found")
	}
	return nil
}

func (r *Repository) CreateAnswer(ctx context.Context, questionID uint, userID uuid.UUID, text string) (*models.Answer, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Question{}).Where("id = ?", questionID).Count(&count).Error; err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("Question not found")
	}

	answer := &models.Answer{
		QuestionID: questionID,
		UserID:     userID,
		Text:       text,
	}

	if err := r.db.WithContext(ctx).Create(answer).Error; err != nil {
		return nil, err
	}
	return answer, nil
}

func (r *Repository) GetAnswerByID(ctx context.Context, id uint) (*models.Answer, error) {
	var answer models.Answer

	if err := r.db.WithContext(ctx).First(&answer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Answer not found")
		}
		return nil, err
	}
	return &answer, nil
}

func (r *Repository) DeleteAnswer(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&models.Answer{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Answer not found")
	}
	return nil
}
