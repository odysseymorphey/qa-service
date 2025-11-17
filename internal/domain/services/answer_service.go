package services

import (
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/repository"
)

type AnswerService interface {
	GetAnswerByID(id int) (*dto.Answer, error)
	CreateAnswer(questionID int, answer dto.Answer) (*dto.Answer, error)
	DeleteAnswer(id int) error
}

type answerService struct {
	repo repository.AnswerRepository
}

func NewAnswerService(repo repository.AnswerRepository) AnswerService {
	return &answerService{
		repo: repo,
	}
}

func (s *answerService) GetAnswerByID(id int) (*dto.Answer, error) {
	return nil, nil
}

func (s *answerService) CreateAnswer(questionID int, answer dto.Answer) (*dto.Answer, error) {
	return nil, nil
}

func (s *answerService) DeleteAnswer(id int) error {
	return nil
}
