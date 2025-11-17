package handlers

import (
	"net/http"
	"qa-service/internal/domain/services"

	"go.uber.org/zap"
)

type QuestionHandler struct {
	service services.QuestionService

	logger *zap.SugaredLogger
}

func NewQuestionHandler(s services.QuestionService, l *zap.SugaredLogger) *QuestionHandler {
	return &QuestionHandler{
		service: s,
		logger:  l,
	}
}

func (h *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {}

func (h *QuestionHandler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {}

func (h *QuestionHandler) DeleteQA(w http.ResponseWriter, r *http.Request) {}
