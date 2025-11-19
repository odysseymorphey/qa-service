package httpapi

import (
	"qa-service/internal/di"
	"qa-service/internal/httpapi/docs"
	"qa-service/internal/httpapi/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(m *chi.Mux, c *di.Container) {
	questionHandler := handlers.NewQuestionHandler(c.GetQuestionService(), c.GetNamedLogger("questionHandler"))
	answerHandler := handlers.NewAnswerHandler(c.GetAnswerService(), c.GetNamedLogger("answerHandler"))

	//	QUESTIONS
	{
		m.Get("/questions", questionHandler.GetAllQuestions)
		m.Get("/questions/{id}", questionHandler.GetQuestionByID)
		m.Post("/questions", questionHandler.CreateQuestion)
		m.Delete("/questions/{id}", questionHandler.DeleteQA)
	}

	// ANSWERS
	{
		m.Get("/answers/{id}", answerHandler.GetAnswerByID)
		m.Post("/questions/{id}/answers", answerHandler.CreateAnswer)
		m.Delete("/answers/{id}", answerHandler.DeleteAnswer)
	}

	m.Get("/docs", docs.ServeIndex)
	m.Get("/docs/swagger.yaml", docs.ServeSwagger)
}
