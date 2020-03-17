package models

import (
	"errors"
	"strings"
	"github.com/badoux/checkmail"
)

type User struct {
	ID       uint32 `json:"id" gorm:"type:bigint;primary_key; AUTO_INCREMENT"`
	Nickname string `json:"nickname "`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index" `
	Password string `json:"password"`
	Name     string `json:"name"`
	Lastname string `json:"string"`
	Deleted  bool    `json:"deleted" gorm:"type:BOOLEAN"`
}
// Validate validates the inputs
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {

	case "login":
		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}

		if u.Password == "" {
			return errors.New("Password is required")
		}
	default:
		if u.Nickname == "" {
			return errors.New("Nickname is required")
		}

		if u.Password == "" {
			return errors.New("Password is required")
		}

		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
	}

	return nil
}