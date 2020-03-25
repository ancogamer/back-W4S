package models

import (
	"strings"
)

type Profile struct {
	ID             uint64 `json:"id" gorm:"type:bigint;primary_key; AUTO_INCREMENT"`
	Avatar         string `json:"avatar" gorm:"type:longtext"` //longtext no BD (mysql-MariaDB)
	DataNascimento string `json:"datanascimento" `             //maximo 8 digitos
	Nickname       string `json:"nickname" `
}

func (p *Profile) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":

	case "login":

	default:

	}
	return nil
}
