package models

import (
	"time"
)

type UserRequest struct {
	ID     int64     `gorm:"column:id;type:int,autoincrement;primaryKey"`
	AccID  int64     `gorm:"column:acc_id;type:int;not null"`
	Status bool      `gorm:"column:status;type:bool;default:false"`
	Date   time.Time `gorm:"column:date;type:time;autoCreateTime:milli"`

	Account *Account `gorm:"foreignKey:acc_id;constraint:OnDelete:CASCADE;"`
}

type UserRequestRepository interface {
	Create(u *UserRequest) error
	Get(id int64) (*UserRequest, error)
	Update(u *UserRequest) error
}
