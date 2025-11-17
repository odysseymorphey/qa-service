package repository

import (
	"qa-service/internal/domain/dto"
	"time"
)

type QuestionRepository interface {
	GetAllQuestions() ([]dto.Question, error)
	GetQuestionByID(id int) (*dto.Question, []dto.Answer, error)
	CreateQuestion(question dto.QuestionRequest) (int, time.Time, error)
	DeleteQA(id int) error
}
