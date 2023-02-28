package repository

import (
	"fmt"
	"gorm.io/gorm"

	"Epam_final/pkg/db/models"
)

type PaymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) Create(row *models.Payments) error {
	err := r.db.Create(&row).Error
	if err != nil {
		return fmt.Errorf("repository create payment: %w", err)
	}

	return nil
}

func (r *PaymentRepo) Get(id int64) (*models.Payments, error) {
	res := &models.Payments{}
	err := r.db.Model(res).
		Preload("Account").
		Where("id = ?", id).
		Find(res).Error
	if err != nil {
		return nil, fmt.Errorf("repository get payment: %w", err)
	}

	return res, nil
}

func (r *PaymentRepo) Update(row *models.Payments) error {
	err := r.db.Updates(&row).
		UpdateColumns("Amount").
		UpdateColumns("Status").
		UpdateColumns("Date").
		Where("user_id = ?", row).Error
	if err != nil {
		return fmt.Errorf("repository update payment: %w", err)
	}

	return nil
}
func (r *PaymentRepo) List(id int64) ([]models.Payments, error) {
	rows := []models.Payments{}
	err := r.db.Model(&rows).
		Preload("Account").
		Where("acc_id = ?", id).
		Find(&rows).Error
	if err != nil {
		return nil, fmt.Errorf("repository accounts list : %w", err)
	}

	return rows, nil
}
