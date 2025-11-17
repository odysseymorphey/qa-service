package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddQuestionsTable, downAddQuestionsTable)
}

func upAddQuestionsTable(ctx context.Context, tx *sql.Tx) error {
	const questionsQuery = `
		CREATE TABLE IF NOT EXISTS questions (
			id BIGSERIAL PRIMARY KEY,
			text TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL
		);
	`

	_, err := tx.ExecContext(ctx, questionsQuery)
	if err != nil {
		return err
	}

	return nil
}

func downAddQuestionsTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE questions;")
	if err != nil {
		return err
	}

	return nil
}
