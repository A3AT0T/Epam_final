package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Epam_final/pkg/db/models"
)

type UserRequestRepo struct {
	db *gorm.DB
}

func NewUserRequestRepo(db *gorm.DB) *UserRequestRepo {
	return &UserRequestRepo{db: db}
}

func (r *UserRequestRepo) Create(row *models.UserRequest) error {
	err := r.db.Create(&row).Error
	if err != nil {
		return fmt.Errorf("repository create UserRequest: %w", err)
	}

	return nil
}

func (r *UserRequestRepo) Get(id int64) (*models.UserRequest, error) {
	res := &models.UserRequest{}
	err := r.db.Model(res).
		Where("id = ?", id).
		Scan(res).Error
	if err != nil {
		return nil, fmt.Errorf("repository get UserRequest: %w", err)
	}

	return res, nil
}

func (r *UserRequestRepo) Update(row *models.UserRequest) error {
	err := r.db.Updates(&row).
		UpdateColumns("status").
		Where("user_id = ?", row).Error
	if err != nil {
		return fmt.Errorf("repository update UserRequest: %w", err)
	}

	return nil
}
