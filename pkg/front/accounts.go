package front

type AccountRes struct {
	ID        int64  `json:"id"`
	Acc       string `json:"acc"`
	UserID    int64  `json:"user_id"`
	AccStatus bool   `json:"acc_status"`
	Amount    int64  `json:"amount"`

	Cards       []CardRes        `json:"cards"`
	UserRequest []UserRequestRes `json:"user_request"`
	Payments    []PaymentsRes    `json:"payments"`
	User        *UserRes         `json:"user"`
}
