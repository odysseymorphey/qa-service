package services

import (
	"errors"
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/repository"
	"strings"
)

var (
	ErrQuestionIDRequired = errors.New("question id is required")
	ErrUserIDRequired     = errors.New("user id is required")
	ErrAnswerTextRequired = errors.New("answer text is required")
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
	return s.repo.GetAnswerByID(id)
}

func (s *answerService) CreateAnswer(questionID int, answer dto.Answer) (*dto.Answer, error) {
	if questionID == 0 {
		return nil, ErrQuestionIDRequired
	}

	if strings.TrimSpace(answer.UserID) == "" {
		return nil, ErrUserIDRequired
	}

	if strings.TrimSpace(answer.Text) == "" {
		return nil, ErrAnswerTextRequired
	}

	answer.QuestionID = questionID

	return s.repo.CreateAnswer(questionID, answer)
}

func (s *answerService) DeleteAnswer(id int) error {
	return s.repo.DeleteAnswer(id)
}
