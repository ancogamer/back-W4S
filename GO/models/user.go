package models

import (
	"errors"
	"fmt"
	"strings"
	"w4s/security"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64  `json:"id" gorm:"type:bigint;primary_key; AUTO_INCREMENT"`
	Nickname  string  `json:"nickname "`
	Email     string  `json:"email" gorm:"type:varchar(100);unique_index" `
	Password  string  `json:"password"`
	Name      string  `json:"name"`
	Lastname  string  `json:"string"`
	Deleted   bool    `json:"deleted" gorm:"type:BOOLEAN"`
	IDProfile uint64  `json:"author_id,omitempty" gorm:"null"`
	Profile   Profile `json:"profile,omitempty"`
	Token     string  `json:"token";sql:"-"`
}

// BeforeSave hash the user password
func (u *User) BeforeSave() error {
	if len(u.Password) > 60 {
		return errors.New("Senha maior que 60 characteres")
	}
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		panic("Password hash")
	}
	u.Password = string(hashedPassword)
	fmt.Println("password salva na struct passada: ", u.Password)
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
