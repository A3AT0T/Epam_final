package repository

import (
	"fmt"

	"gorm.io/gorm"

	"Epam_final/pkg/db/models"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) Create(row *models.Account) error {
	err := r.db.Create(&row).Error
	if err != nil {
		return fmt.Errorf("repository create account: %w", err)
	}

	return nil
}

func (r *AccountRepo) Get(id int64) (*models.Account, error) {
	res := &models.Account{}
	err := r.db.Model(res).
		Preload("cards").
		Preload("unlock_table").
		Preload("payments").
		Preload("user").
		Where("id = ?", id).
		Scan(res).Error
	if err != nil {
		return nil, fmt.Errorf("repository get account: %w", err)
	}

	return res, nil
}

func (r *AccountRepo) List(id int64) ([]models.Account, error) {
	var rows []models.Account
	err := r.db.Model(rows).
		Preload("cards").
		Preload("unlock_table").
		Preload("payments").
		Preload("user").
		Where("user_id = ?", id).
		Scan(rows).Error
	if err != nil {
		return nil, fmt.Errorf("repository accounts list : %w", err)
	}

	return rows, nil
}

func (r *AccountRepo) Update(row *models.Account) error {
	err := r.db.Updates(&row).
		UpdateColumns("amount").
		UpdateColumns("acc_status").Error
	if err != nil {
		return fmt.Errorf("repository update account: %w", err)
	}

	return nil
}
