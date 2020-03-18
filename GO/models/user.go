package models

import (
	"errors"
	"strings"
	"github.com/badoux/checkmail"
	"w4s/security"
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
// BeforeSave hash the user password
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
// Validate validates the inputs
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Nickname is required")
		}

		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
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