package di

import (
	"log"
	"qa-service/internal/config"
	"qa-service/internal/database/migrator"
	"qa-service/internal/database/postgres"
	"qa-service/internal/database/repository"
	"qa-service/internal/domain/services"

	"go.uber.org/zap"
)

type Container struct {
	answerService   services.AnswerService
	questionService services.QuestionService

	logger *zap.Logger
}

func NewContainer(cfg *config.Config) *Container {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	db, err := postgres.New(cfg.DBConfig.DSN())
	if err != nil {
		log.Fatalf("can't open database: %v", err)
	}

	if err := migrator.Run(db); err != nil {
		log.Fatalf("can't apply migrations: %v", err)
	}

	answerRepo := repository.NewAnswerRepo(db)
	questionRepo := repository.NewQuestionRepo(db)

	answerService := services.NewAnswerService(answerRepo)
	questionService := services.NewQuestionService(questionRepo)

	return &Container{
		answerService:   answerService,
		questionService: questionService,
		logger:          zapLogger,
	}
}

func (c *Container) GetAnswerService() services.AnswerService {
	return c.answerService
}

func (c *Container) GetQuestionService() services.QuestionService {
	return c.questionService
}

func (c *Container) GetNamedLogger(name string) *zap.SugaredLogger {
	return c.logger.Named(name).Sugar()
}
