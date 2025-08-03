package dto

type EmailCode struct {
	Email string `json:"email" binding:"required,email"`
}
