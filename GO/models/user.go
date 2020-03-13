package models
//model do struct

type User struct {
	ID       uint32 `json:"id" gorm:"type:bigint;primary_key; AUTO_INCREMENT"`
	Nickname string `json:"nickname "`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index" `
	Password string `json:"password"`
	Name     string `json:"name"`
	Lastname string `json:"string"`
	Deleted  bool    `json:"deleted" gorm:"type:BOOLEAN"`
}
