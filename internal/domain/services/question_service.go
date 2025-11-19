package services

import (
	"errors"
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/repository"
	"strings"
)

var ErrQuestionTextRequired = errors.New("question text is required")

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
	return s.repo.GetAllQuestions()
}

func (s *questionService) GetQuestionByID(id int) (*dto.Question, []dto.Answer, error) {
	return s.repo.GetQuestionByID(id)
}

func (s *questionService) CreateQuestion(question dto.QuestionRequest) (*dto.Question, error) {
	if strings.TrimSpace(question.Text) == "" {
		return nil, ErrQuestionTextRequired
	}

	id, createdAt, err := s.repo.CreateQuestion(question)
	if err != nil {
		return nil, err
	}

	return &dto.Question{
		ID:        id,
		Text:      question.Text,
		CreatedAt: createdAt,
	}, nil
}

func (s *questionService) DeleteQA(id int) error {
	return s.repo.DeleteQA(id)
}
