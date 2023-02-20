package front

type CardRes struct {
	ID     int64  `json:"id"`
	CardID string `json:"card_id"`
	AccID  int64  `json:"acc_id"`
}
