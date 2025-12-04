package models

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Text      string    `json:"text" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	Answers   []Answer  `json:"answers,omitempty" gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE;"`
}

type Answer struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	QuestionID uint      `json:"question_id" gorm:"not null;index"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Text       string    `json:"text" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateQuestionRequest struct {
	Text string `json:"text" validate:"required"`
}

type CreateAnswerRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Text   string    `json:"text" validate:"required"`
}
