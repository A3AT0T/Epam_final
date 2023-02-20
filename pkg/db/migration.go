package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Epam_final/pkg/db/models"
)

func GetConn(address string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(address))
	if err != nil {
		return nil, fmt.Errorf("connect to db: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	if err := db.Migrator().
		AutoMigrate(
			&models.Account{},
			&models.Payments{},
			&models.UserRequest{},
			&models.User{},
			&models.Card{},
			&models.Log{},
		); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
