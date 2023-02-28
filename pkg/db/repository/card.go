package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Epam_final/pkg/db/models"
)

type CardRepo struct {
	db *gorm.DB
}

func NewCardRepo(db *gorm.DB) *CardRepo {
	return &CardRepo{db: db}
}

func (r *CardRepo) Create(row *models.Card) error {
	err := r.db.Create(&row).Error
	if err != nil {
		return fmt.Errorf("repository create card: %w", err)
	}

	return nil
}

func (r *CardRepo) Get(id int64) (*models.Card, error) {
	res := &models.Card{}
	err := r.db.Model(res).
		Preload("Account").
		Where("id = ?", id).
		Find(res).Error
	if err != nil {
		return nil, fmt.Errorf("repository get card: %w", err)
	}

	return res, nil
}

func (r *CardRepo) List(id int64) ([]models.Card, error) {
	rows := []models.Card{}
	err := r.db.Model(rows).
		Preload("Account").
		Where("user_id = ?", id).
		Find(rows).Error
	if err != nil {
		return nil, fmt.Errorf("repository account list : %w", err)
	}

	return rows, nil
}
