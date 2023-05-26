package model

type User struct {
	ID 			uint64 	`json:"id" form:"id" gorm:"primarykey;not null"`	
	Email       string    `json:"email" form:"email" gorm:"unique;not null"`
	Password 	string     `json:"password" form:"password"`
	Token       string `json:"-" form:"-"`
}

