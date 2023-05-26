package model

type Item struct{
	ID string `json:"id" form:"id" gorm:"primarykey"`
	Name string `json:"name" form:"name" gorm:"not null"`
	CategoryID uint64 `json:"category_id" gorm"not null"`
}