package repository

import (
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/repository"
	"time"

	"gorm.io/gorm"
)

type questionsRepo struct {
	db *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) repository.QuestionRepository {
	return &questionsRepo{
		db: db,
	}
}

func (r *questionsRepo) GetAllQuestions() ([]dto.Question, error) {
	return nil, nil
}

func (r *questionsRepo) GetQuestionByID(id int) (*dto.Question, []dto.Answer, error) {
	return nil, nil, nil
}
func (r *questionsRepo) CreateQuestion(question dto.QuestionRequest) (int, time.Time, error) {
	return nil, nil
}
func (r *questionsRepo) DeleteQA(id int) error {
	return nil
}
