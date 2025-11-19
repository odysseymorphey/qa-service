package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/services"

	"go.uber.org/zap"
	"gorm.io/gorm"
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

func (h *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.service.GetAllQuestions()
	if err != nil {
		h.logger.Errorf("failed to get questions: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to load questions")
		return
	}

	writeJSON(w, http.StatusOK, questions)
}

func (h *QuestionHandler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	question, answers, err := h.service.GetQuestionByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "question not found")
			return
		}

		h.logger.Errorf("failed to get question by id %d: %v", id, err)
		writeError(w, http.StatusInternalServerError, "failed to load question")
		return
	}

	response := struct {
		dto.Question
		Answers []dto.Answer `json:"answers"`
	}{
		Question: *question,
		Answers:  answers,
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req dto.QuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	question, err := h.service.CreateQuestion(req)
	if err != nil {
		if errors.Is(err, services.ErrQuestionTextRequired) {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		h.logger.Errorf("failed to create question: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to create question")
		return
	}

	writeJSON(w, http.StatusCreated, question)
}

func (h *QuestionHandler) DeleteQA(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.DeleteQA(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "question not found")
			return
		}

		h.logger.Errorf("failed to delete question %d: %v", id, err)
		writeError(w, http.StatusInternalServerError, "failed to delete question")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
