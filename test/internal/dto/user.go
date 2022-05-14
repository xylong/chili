package dto

type LoginParam struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=6,max=18"`
}
