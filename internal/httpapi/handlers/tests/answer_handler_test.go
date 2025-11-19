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

type mockAnswerService struct {
	getFunc    func(id int) (*dto.Answer, error)
	createFunc func(questionID int, answer dto.Answer) (*dto.Answer, error)
	deleteFunc func(id int) error
}

func (m *mockAnswerService) GetAnswerByID(id int) (*dto.Answer, error) {
	if m.getFunc != nil {
		return m.getFunc(id)
	}

	return nil, nil
}

func (m *mockAnswerService) CreateAnswer(questionID int, answer dto.Answer) (*dto.Answer, error) {
	if m.createFunc != nil {
		return m.createFunc(questionID, answer)
	}

	return nil, nil
}

func (m *mockAnswerService) DeleteAnswer(id int) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(id)
	}

	return nil
}

func TestAnswerHandler_GetAnswerByID_Success(t *testing.T) {
	mockService := &mockAnswerService{
		getFunc: func(id int) (*dto.Answer, error) {
			return &dto.Answer{ID: id, Text: "hello"}, nil
		},
	}

	handler := handlers.NewAnswerHandler(mockService, zap.NewNop().Sugar())

	req := withChiURLParam(httptest.NewRequest(http.MethodGet, "/answers/4", nil), "id", "4")
	rr := httptest.NewRecorder()

	handler.GetAnswerByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var ans dto.Answer
	if err := json.Unmarshal(rr.Body.Bytes(), &ans); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if ans.ID != 4 {
		t.Fatalf("unexpected response: %+v", ans)
	}
}

func TestAnswerHandler_GetAnswerByID_NotFound(t *testing.T) {
	mockService := &mockAnswerService{
		getFunc: func(id int) (*dto.Answer, error) {
			return nil, gorm.ErrRecordNotFound
		},
	}

	handler := handlers.NewAnswerHandler(mockService, zap.NewNop().Sugar())

	req := withChiURLParam(httptest.NewRequest(http.MethodGet, "/answers/2", nil), "id", "2")
	rr := httptest.NewRecorder()

	handler.GetAnswerByID(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}
}

func TestAnswerHandler_CreateAnswer_Validation(t *testing.T) {
	mockService := &mockAnswerService{
		createFunc: func(questionID int, answer dto.Answer) (*dto.Answer, error) {
			return nil, services.ErrAnswerTextRequired
		},
	}

	handler := handlers.NewAnswerHandler(mockService, zap.NewNop().Sugar())

	body := bytes.NewBufferString(`{"user_id":"","text":""}`)
	req := withChiURLParam(httptest.NewRequest(http.MethodPost, "/questions/1/answers", body), "id", "1")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateAnswer(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
}

func TestAnswerHandler_CreateAnswer_InternalError(t *testing.T) {
	mockService := &mockAnswerService{
		createFunc: func(questionID int, answer dto.Answer) (*dto.Answer, error) {
			return nil, errors.New("boom")
		},
	}

	handler := handlers.NewAnswerHandler(mockService, zap.NewNop().Sugar())

	body := bytes.NewBufferString(`{"user_id":"u1","text":"hi"}`)
	req := withChiURLParam(httptest.NewRequest(http.MethodPost, "/questions/2/answers", body), "id", "2")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateAnswer(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
}

func TestAnswerHandler_CreateAnswer_Success(t *testing.T) {
	mockService := &mockAnswerService{
		createFunc: func(questionID int, answer dto.Answer) (*dto.Answer, error) {
			return &dto.Answer{ID: 7, QuestionID: questionID, Text: answer.Text, UserID: answer.UserID}, nil
		},
	}

	handler := handlers.NewAnswerHandler(mockService, zap.NewNop().Sugar())

	body := bytes.NewBufferString(`{"user_id":"user","text":"answer"}`)
	req := withChiURLParam(httptest.NewRequest(http.MethodPost, "/questions/3/answers", body), "id", "3")
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateAnswer(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}

	var ans dto.Answer
	if err := json.Unmarshal(rr.Body.Bytes(), &ans); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if ans.ID != 7 || ans.QuestionID != 3 || ans.Text != "answer" {
		t.Fatalf("unexpected response: %+v", ans)
	}
}

func TestAnswerHandler_DeleteAnswer_NotFound(t *testing.T) {
	mockService := &mockAnswerService{
		deleteFunc: func(id int) error {
			return gorm.ErrRecordNotFound
		},
	}

	handler := handlers.NewAnswerHandler(mockService, zap.NewNop().Sugar())

	req := withChiURLParam(httptest.NewRequest(http.MethodDelete, "/answers/9", nil), "id", "9")
	rr := httptest.NewRecorder()

	handler.DeleteAnswer(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}
}

func TestAnswerHandler_DeleteAnswer_Success(t *testing.T) {
	mockService := &mockAnswerService{
		deleteFunc: func(id int) error {
			if id != 11 {
				t.Fatalf("unexpected id: %d", id)
			}
			return nil
		},
	}

	handler := handlers.NewAnswerHandler(mockService, zap.NewNop().Sugar())

	req := withChiURLParam(httptest.NewRequest(http.MethodDelete, "/answers/11", nil), "id", "11")
	rr := httptest.NewRecorder()

	handler.DeleteAnswer(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", rr.Code)
	}
}
