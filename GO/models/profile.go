package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"strings"
)

type Profile struct {
	ID             uint64  `json:"id" gorm:"type:bigint;primary_key; AUTO_INCREMENT"`
	Avatar         string  `json:"avatar" gorm:"type:longtext"` //longtext no BD (mysql-MariaDB)
	DataNascimento string  `json:"datanascimento" `//maximo 8 digitos
	Nickname       string  `json:"nickname" `
}
func (p *Profile) Validate(action string) error {
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