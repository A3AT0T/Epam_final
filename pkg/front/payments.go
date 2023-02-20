package front

import (
	"fmt"
	"time"
)

type PaymentsRes struct {
	ID     int64     `json:"id"`
	AccID  int64     `json:"acc_id"`
	Amount int64     `json:"amount"`
	Status bool      `json:"status"`
	Date   time.Time `json:"date"`

	Account *AccountRes `json:"account"`
}

type PaymentReq struct {
	AccID  int64 `json:"acc_id"`
	Amount int64 `json:"amount"`
	Status bool  `json:"status"`
}

func (p *PaymentReq) Validate() error {
	if p.Amount == 0 {
		return fmt.Errorf("transaction can't be zero")
	}
	if p.AccID == 0 {
		return fmt.Errorf("accountID can't be zero")
	}
	return nil
}
