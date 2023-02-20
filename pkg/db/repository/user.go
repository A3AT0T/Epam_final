package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Epam_final/pkg/db/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(row *models.User) error {
	err := r.db.Create(&row).Error
	if err != nil {
		return fmt.Errorf("repository create user: %w", err)
	}

	return nil
}

func (r *UserRepo) Get(id int64) (*models.User, error) {
	res := &models.User{}
	err := r.db.Model(res).
		Preload("accounts").
		Preload("logs").
		Where("id = ?", id).
		Scan(res).Error
	if err != nil {
		return nil, fmt.Errorf("repository get user: %w", err)
	}

	return res, nil
}

func (r *UserRepo) Update(row *models.User) error {
	err := r.db.Updates(&row).
		UpdateColumns("status").
		Where("id = ?", row.ID).Error
	if err != nil {
		return fmt.Errorf("repository update user: %w", err)
	}

	return nil
}
