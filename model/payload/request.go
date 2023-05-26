package payload

type UploadItemRequest struct {
	Name     string `json:"name" form:"name" gorm:"not null"`
	CategoryID uint64 `json:"category_id" gorm"not null"`
}

type CreateItemRequest struct {
	Name     string `json:"name" form:"name" gorm:"not null"`
	CategoryID uint64 `json:"category_id" gorm"not null"`
}

type UpdateItemRequest struct {
	Name     string `json:"name" form:"name" gorm:"not null"`
	CategoryID uint64 `json:"category_id" gorm"not null"`
}

type UserRegisterRequst struct {
	Email       string    `json:"email" form:"email" gorm:"unique;not null"`
	Password 	string     `json:"password" form:"password"`
}

type UserLoginRequest struct {
	Email       string    `json:"email" form:"email" gorm:"unique;not null"`
	Password 	string     `json:"password" form:"password"`
}