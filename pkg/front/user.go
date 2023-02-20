package front

type UserRes struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Pass    string `json:"pass"`
	Status  string `json:"status"`
	IsAdmin bool   `json:"is_admin"`

	Accounts []AccountRes `json:"accounts"`
	Logs     []LogRes     `json:"logs"`
}

type UserReq struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Pass    string `json:"pass"`
}
