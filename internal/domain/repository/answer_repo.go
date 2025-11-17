package repository

import "qa-service/internal/domain/dto"

type AnswerRepository interface {
	GetAnswerByID(id int) (*dto.Answer, error)
	CreateAnswer(questionID int, answer dto.Answer) (*dto.Answer, error)
	DeleteAnswer(id int) error
}
