package front

import "time"

type UserRequestRes struct {
	ID    int64 `json:"id"`
	AccID int64 `json:"acc_id"`
	// status false=new request, true=completed/approved
	Status bool      `json:"status"`
	Date   time.Time `json:"date"`
}

type UserRequest struct {
	AccID int64 `json:"acc_id"`
}
