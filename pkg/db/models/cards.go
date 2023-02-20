package models

type Card struct {
	ID     int64  `gorm:"column:id;type:int;primaryKey;autoincrement"`
	CardID string `gorm:"column:card_id;type:string;not null"`
	AccID  int64  `gorm:"column:acc_id;type:int;not null"`

	Account *Account `gorm:"foreignKey:acc_id"`
}

type CardsRepository interface {
	Create(c *Card) error
	Get(id int64) (*Card, error)
}
