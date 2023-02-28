package models

import (
	"time"
)

type Log struct {
	ID      int64     `gorm:"column:id;type:int;autoincrement;primaryKey"`
	UserID  int64     `gorm:"column:user_id;type:int;not null"`
	Massage string    `gorm:"column:massage;type:string"`
	Date    time.Time `gorm:"column:date;type:time;autoCreateTime:milli"`
}

type LogsRepository interface {
	Create(l *Log) error
	Get(userID int64) (*Log, error)
	List(userID int64) ([]Log, error)
}
