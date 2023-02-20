package models

import (
	"time"
)

type Payments struct {
	ID     int64     `gorm:"column:id;type:int;primaryKey;autoincrement"`
	AccID  int64     `gorm:"column:acc_id;type:int;not null"`
	Amount int64     `gorm:"column:amount;type:int;not null"`
	Status bool      `gorm:"column:status;type:bool;default:false"`
	Date   time.Time `gorm:"column:date;type:time;autoCreateTime:milli"`

	Account *Account `gorm:"foreignKey:acc_id"`
}

type PaymentsRepository interface {
	Create(p *Payments) error
	Get(id int64) (*Payments, error)
	Update(p *Payments) error
	List(accID int64) ([]Payments, error)
}
