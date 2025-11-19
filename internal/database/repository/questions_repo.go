package repository

import (
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/repository"
	"time"

	"gorm.io/gorm"
)

type questionsRepo struct {
	db *gorm.DB
}

func NewQuestionRepo(db *gorm.DB) repository.QuestionRepository {
	return &questionsRepo{
		db: db,
	}
}

func (r *questionsRepo) GetAllQuestions() ([]dto.Question, error) {
	var questions []dto.Question

	if err := r.db.Table("questions").
		Select("id, text, created_at").
		Order("id").
		Scan(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *questionsRepo) GetQuestionByID(id int) (*dto.Question, []dto.Answer, error) {
	var question dto.Question

	if err := r.db.Table("questions").
		Where("id = ?", id).
		Select("id, text, created_at").
		Take(&question).Error; err != nil {
		return nil, nil, err
	}

	var answers []dto.Answer

	if err := r.db.Table("answers").
		Where("question_id = ?", id).
		Select("id, question_id, user_id, text, created_at").
		Order("id").
		Scan(&answers).Error; err != nil {
		return nil, nil, err
	}

	return &question, answers, nil
}
func (r *questionsRepo) CreateQuestion(question dto.QuestionRequest) (int, time.Time, error) {
	now := time.Now().UTC()

	type createResult struct {
		ID        int
		CreatedAt time.Time
	}

	var result createResult

	err := r.db.Raw(`
		INSERT INTO questions (text, created_at)
		VALUES (?, ?)
		RETURNING id, created_at
	`, question.Text, now).Scan(&result).Error
	if err != nil {
		return 0, time.Time{}, err
	}

	return result.ID, result.CreatedAt, nil
}
func (r *questionsRepo) DeleteQA(id int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("answers").Where("question_id = ?", id).Delete(&dto.Answer{}).Error; err != nil {
			return err
		}

		result := tx.Table("questions").Where("id = ?", id).Delete(&dto.Question{})
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil
	})
}
