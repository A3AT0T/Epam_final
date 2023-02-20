package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Epam_final/pkg/db/models"
)

type LogRepo struct {
	db *gorm.DB
}

func NewLogRepo(db *gorm.DB) *LogRepo {
	return &LogRepo{db: db}
}

func (r *LogRepo) Create(row *models.Log) error {
	err := r.db.Create(&row).Error
	if err != nil {
		return fmt.Errorf("repository create log: %w", err)
	}

	return nil
}

func (r *LogRepo) Get(id int64) (*models.Log, error) {
	rows := &models.Log{}
	err := r.db.Model(rows).
		Where("user_id = ?", id).
		Scan(rows).Error
	if err != nil {
		return nil, fmt.Errorf("repository log : %w", err)
	}

	return rows, nil
}

func (r *LogRepo) List(id int64) ([]models.Log, error) {
	rows := []models.Log{}
	err := r.db.Model(rows).
		Where("user_id = ?", id).
		Scan(rows).Error
	if err != nil {
		return nil, fmt.Errorf("repository logs list : %w", err)
	}

	return rows, nil
}
