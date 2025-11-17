package handlers

import (
	"net/http"
	"qa-service/internal/domain/services"

	"go.uber.org/zap"
)

type AnswerHandler struct {
	service services.AnswerService

	logger *zap.SugaredLogger
}

func NewAnswerHandler(s services.AnswerService, l *zap.SugaredLogger) *AnswerHandler {
	return &AnswerHandler{
		service: s,
		logger:  l,
	}
}

func (h *AnswerHandler) GetAnswerByID(w http.ResponseWriter, r *http.Request) {

}

func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {}

func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {}
