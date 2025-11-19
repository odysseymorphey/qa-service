package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/services"
	"qa-service/internal/httpapi/handlers"
)

type mockQuestionService struct {
	getAllFunc   func() ([]dto.Question, error)
	getByIDFunc  func(id int) (*dto.Question, []dto.Answer, error)
	createFunc   func(question dto.QuestionRequest) (*dto.Question, error)
	deleteQAfunc func(id int) error
}

func (m *mockQuestionService) GetAllQuestions() ([]dto.Question, error) {
	if m.getAllFunc != nil {
		return m.getAllFunc()
	}

	return nil, nil
}

func (m *mockQuestionService) GetQuestionByID(id int) (*dto.Question, []dto.Answer, error) {
	if m.getByIDFunc != nil {
		return m.getByIDFunc(id)
	}

	return nil, nil, nil
}

func (m *mockQuestionService) CreateQuestion(question dto.QuestionRequest) (*dto.Question, error) {
	if m.createFunc != nil {
		return m.createFunc(question)
	}

	return nil, nil
}

func (m *mockQuestionService) DeleteQA(id int) error {
	if m.deleteQAfunc != nil {
		return m.deleteQAfunc(id)
	}

	return nil
}

func TestQuestionHandler_GetAllQuestions_Success(t *testing.T) {
	mockService := &mockQuestionService{
		getAllFunc: func() ([]dto.Question, error) {
			return []dto.Question{{ID: 1, Text: "test"}}, nil
		},
	}

	handler := handlers.NewQuestionHandler(mockService, zap.NewNop().Sugar())

	req := httptest.NewRequest(http.MethodGet, "/questions", nil)
	rr := httptest.NewRecorder()

	handler.GetAllQuestions(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp []dto.Question
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(resp) != 1 || resp[0].ID != 1 {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestQuestionHandler_GetAllQuestions_Error(t *testing.T) {
	mockService := &mockQuestionService{
		getAllFunc: func() ([]dto.Question, error) {
			return nil, errors.New("boom")
		},
	}

	handler := handlers.NewQuestionHandler(mockService, zap.NewNop().Sugar())

	req := httptest.NewRequest(http.MethodGet, "/questions", nil)
	rr := httptest.NewRecorder()

	handler.GetAllQuestions(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
}

func TestQuestionHandler_GetQuestionByID_NotFound(t *testing.T) {
	mockService := &mockQuestionService{
		getByIDFunc: func(id int) (*dto.Question, []dto.Answer, error) {
			return nil, nil, gorm.ErrRecordNotFound
		},
	}

	handler := handlers.NewQuestionHandler(mockService, zap.NewNop().Sugar())

	req := withChiURLParam(httptest.NewRequest(http.MethodGet, "/questions/1", nil), "id", "1")
	rr := httptest.NewRecorder()

	handler.GetQuestionByID(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}
}

func TestQuestionHandler_CreateQuestion_Validation(t *testing.T) {
	mockService := &mockQuestionService{
		createFunc: func(question dto.QuestionRequest) (*dto.Question, error) {
			return nil, services.ErrQuestionTextRequired
		},
	}

	handler := handlers.NewQuestionHandler(mockService, zap.NewNop().Sugar())

	body := bytes.NewBufferString(`{"text":""}`)
	req := httptest.NewRequest(http.MethodPost, "/questions", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateQuestion(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestQuestionHandler_CreateQuestion_Success(t *testing.T) {
	mockService := &mockQuestionService{
		createFunc: func(question dto.QuestionRequest) (*dto.Question, error) {
			return &dto.Question{ID: 10, Text: question.Text}, nil
		},
	}

	handler := handlers.NewQuestionHandler(mockService, zap.NewNop().Sugar())

	body := bytes.NewBufferString(`{"text":"hello"}`)
	req := httptest.NewRequest(http.MethodPost, "/questions", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateQuestion(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}

	var resp dto.Question
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.ID != 10 || resp.Text != "hello" {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestQuestionHandler_DeleteQA_NotFound(t *testing.T) {
	mockService := &mockQuestionService{
		deleteQAfunc: func(id int) error {
			return gorm.ErrRecordNotFound
		},
	}

	handler := handlers.NewQuestionHandler(mockService, zap.NewNop().Sugar())

	req := withChiURLParam(httptest.NewRequest(http.MethodDelete, "/questions/5", nil), "id", "5")
	rr := httptest.NewRecorder()

	handler.DeleteQA(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}
}

func TestQuestionHandler_DeleteQA_Success(t *testing.T) {
	mockService := &mockQuestionService{
		deleteQAfunc: func(id int) error {
			if id != 3 {
				t.Fatalf("unexpected id passed: %d", id)
			}
			return nil
		},
	}

	handler := handlers.NewQuestionHandler(mockService, zap.NewNop().Sugar())

	req := withChiURLParam(httptest.NewRequest(http.MethodDelete, "/questions/3", nil), "id", "3")
	rr := httptest.NewRecorder()

	handler.DeleteQA(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", rr.Code)
	}
}
