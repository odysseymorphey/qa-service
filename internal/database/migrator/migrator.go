package migrator

import (
	"context"
	"time"

	"github.com/pressly/goose/v3"
	"gorm.io/gorm"

	_ "qa-service/migrations"
)

const (
	migrationsDir    = "migrations"
	migrationTimeout = 30 * time.Second
)

func Run(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), migrationTimeout)
	defer cancel()

	return goose.UpContext(ctx, sqlDB, migrationsDir)
}
