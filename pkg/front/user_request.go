package front

import "time"

type UserRequestRes struct {
	ID     int64     `json:"id"`
	AccID  int64     `json:"acc_id"`
	Status bool      `json:"status"`
	Date   time.Time `json:"date"`
}
