package repository

import (
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/repository"

	"gorm.io/gorm"
)

type answerRepo struct {
	db *gorm.DB
}

func NewAnswerRepo(db *gorm.DB) repository.AnswerRepository {
	return &answerRepo{
		db: db,
	}
}

func (r *answerRepo) GetAnswerByID(id int) (*dto.Answer, error) {
	return nil, nil
}
func (r *answerRepo) CreateAnswer(questionID int, answer dto.Answer) (*dto.Answer, error) {
	return nil, nil
}

func (r *answerRepo) DeleteAnswer(id int) error {
	return nil
}
