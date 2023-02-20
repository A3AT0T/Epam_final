package front

import (
	"time"
)

type LogRes struct {
	ID      int64     `json:"id"`
	UserID  int64     `json:"user_id"`
	Massage string    `json:"massage"`
	Date    time.Time `json:"date"`
}
