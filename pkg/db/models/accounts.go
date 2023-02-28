package models

type Account struct {
	ID        int64  `gorm:"column:id;type:int;autoincrement;primaryKey"`
	Acc       string `gorm:"column:acc;type:string;size:31;unique;not null"`
	UserID    int64  `gorm:"column:user_id;type:int;not null"`
	AccStatus bool   `gorm:"column:acc_status;type:bool;default:true;not null"`
	Amount    int64  `gorm:"column:amount;type:int;default:0"`

	Cards       []Card        `gorm:"foreignKey:acc_id;references:id;constraint:OnDelete:CASCADE;"`
	UserRequest []UserRequest `gorm:"foreignKey:acc_id;references:id;constraint:OnDelete:CASCADE;"`
	Payments    []Payments    `gorm:"foreignKey:acc_id;references:id;constraint:OnDelete:CASCADE;"`
	User        *User         `gorm:"foreignKey:user_id;constraint:OnDelete:CASCADE;"`
}

type AccountsRepository interface {
	Create(a *Account) error
	Get(id int64) (*Account, error)
	List(userID int64) ([]Account, error)
	Update(a *Account) error
}
