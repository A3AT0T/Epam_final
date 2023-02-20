package models

type User struct {
	ID      int64  `gorm:"column:id;type:int;primaryKey;autoincrement"`
	Name    string `gorm:"column:name;type:string;not null"`
	Surname string `gorm:"column:surname;type:string;not null"`
	Email   string `gorm:"column:email;type:string;not null;unique"`
	Pass    string `gorm:"column:pass;type:string;not null"`
	Status  string `gorm:"column:status;type:string;not null;default:active"`
	IsAdmin bool   `gorm:"column:is_admin;type:bool;not null;default:false"`

	Accounts []Account `gorm:"foreignKey:user_id;references:id"`
	Logs     []Log     `gorm:"foreignKey:user_id;references:id"`
}

type UserRepository interface {
	Create(u *User) error
	Get(id int64) (*User, error)
	Update(u *User) error
}
