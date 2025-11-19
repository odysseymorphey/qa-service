package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/services"

	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	id, err := parseIDParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	answer, err := h.service.GetAnswerByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "answer not found")
			return
		}

		h.logger.Errorf("failed to get answer %d: %v", id, err)
		writeError(w, http.StatusInternalServerError, "failed to load answer")
		return
	}

	writeJSON(w, http.StatusOK, answer)
}

func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	questionID, err := parseIDParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var req dto.AnswerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	answer, err := h.service.CreateAnswer(questionID, dto.Answer{
		UserID: req.UserID,
		Text:   req.Text,
	})
	if err != nil {
		switch {
		case errors.Is(err, services.ErrQuestionIDRequired),
			errors.Is(err, services.ErrUserIDRequired),
			errors.Is(err, services.ErrAnswerTextRequired):
			writeError(w, http.StatusBadRequest, err.Error())
			return
		default:
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23503" {
				writeError(w, http.StatusBadRequest, "question not found")
				return
			}

			h.logger.Errorf("failed to create answer: %v", err)
			writeError(w, http.StatusInternalServerError, "failed to create answer")
			return
		}
	}

	writeJSON(w, http.StatusCreated, answer)
}

func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeleteAnswer(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "answer not found")
			return
		}

		h.logger.Errorf("failed to delete answer %d: %v", id, err)
		writeError(w, http.StatusInternalServerError, "failed to delete answer")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
