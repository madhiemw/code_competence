package payload

type GetAllItemResponse struct {
	ID         string `json:"id" form:"id" gorm:"primarykey"`
	Name       string `json:"name" form:"name" gorm:"not null"`
	CategoryID uint64 `json:"category_id" gorm"not null"`
}

type UpdateItemResponse struct {
	Name       string `json:"name" form:"name" gorm:"not null"`
	CategoryID uint64 `json:"category_id" gorm"not null"`
}

type UserRegisterResponse struct {
	Email    string `json:"email" form:"email" gorm:"unique;not null"`
	Password string `json:"password" form:"password"`
}
type LoginUserResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
