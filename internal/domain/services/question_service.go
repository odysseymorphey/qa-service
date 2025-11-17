package services

import (
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/repository"
)

type QuestionService interface {
	GetAllQuestions() ([]dto.Question, error)
	GetQuestionByID(id int) (*dto.Question, []dto.Answer, error)
	CreateQuestion(question dto.QuestionRequest) (*dto.Question, error)
	DeleteQA(id int) error
}

type questionService struct {
	repo repository.QuestionRepository
}

func NewQuestionService(repo repository.QuestionRepository) QuestionService {
	return &questionService{
		repo: repo,
	}
}

func (s *questionService) GetAllQuestions() ([]dto.Question, error) {
	return nil, nil
}

func (s *questionService) GetQuestionByID(id int) (*dto.Question, []dto.Answer, error) {
	return nil, nil, nil
}

func (s *questionService) CreateQuestion(question dto.QuestionRequest) (*dto.Question, error) {
	return nil, nil
}

func (s *questionService) DeleteQA(id int) error {
	return nil
}
