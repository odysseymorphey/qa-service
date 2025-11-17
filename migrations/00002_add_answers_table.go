package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddAnswersTable, downAddAnswersTable)
}

func upAddAnswersTable(ctx context.Context, tx *sql.Tx) error {
	const answersQuery = `
		CREATE TABLE IF NOT EXISTS answers (
			id BIGSERIAL PRIMARY KEY,
			question_id BIGINT NOT NULL REFERENCES questions(id),
			user_id VARCHAR(50) NOT NULL,
			text TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL
		);
	`

	_, err := tx.ExecContext(ctx, answersQuery)
	if err != nil {
		return err
	}

	return nil
}

func downAddAnswersTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE answers")
	if err != nil {
		return err
	}

	return nil
}
