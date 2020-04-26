package models

type LoginUser struct {
	Email    string `json:"email" `
	Nickname string `json:"nickname" `
	Password string `json:"password" binding:"required"`
	Token    string `json:"token"`
}
