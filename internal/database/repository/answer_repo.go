package repository

import (
	"qa-service/internal/domain/dto"
	"qa-service/internal/domain/repository"
	"time"

	"gorm.io/gorm"
)

type answerRepo struct {
	db *gorm.DB
}

func NewAnswerRepo(db *gorm.DB) repository.AnswerRepository {
	return &answerRepo{
		db: db,
	}
}

func (r *answerRepo) GetAnswerByID(id int) (*dto.Answer, error) {
	var answer dto.Answer

	if err := r.db.Table("answers").Where("id = ?", id).Take(&answer).Error; err != nil {
		return nil, err
	}

	return &answer, nil
}
func (r *answerRepo) CreateAnswer(questionID int, answer dto.Answer) (*dto.Answer, error) {
	now := time.Now().UTC()

	type createdAnswer struct {
		ID         int
		QuestionID int
		UserID     string
		Text       string
		CreatedAt  time.Time
	}

	var created createdAnswer

	err := r.db.Raw(`
		INSERT INTO answers (question_id, user_id, text, created_at)
		VALUES (?, ?, ?, ?)
		RETURNING id, question_id, user_id, text, created_at
	`, questionID, answer.UserID, answer.Text, now).Scan(&created).Error
	if err != nil {
		return nil, err
	}

	return &dto.Answer{
		ID:         created.ID,
		QuestionID: created.QuestionID,
		UserID:     created.UserID,
		Text:       created.Text,
		CreatedAt:  created.CreatedAt,
	}, nil
}

func (r *answerRepo) DeleteAnswer(id int) error {
	result := r.db.Table("answers").Where("id = ?", id).Delete(&dto.Answer{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
