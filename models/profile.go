package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type Profile struct {
	gorm.Model
	IDUser         uint   `json:"id_user" gorm:"default:0"`
	Nickname       string `json:"nickname" gorm:"unique_index"` //max 15
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	Avatar         string `json:"avatar"`          //longtext no BD (mysql-MariaDB)
	DataNascimento string `json:"datanascimento" ` //maximo 8 digitos
	Deleted        bool   `json:"deleted" gorm:"type:BOOLEAN"`
}
type ProfileInput struct {
	Nickname       string `json:"nickname" binding:"required"`
	Avatar         string `json:"avatar"`          //longtext no BD (mysql-MariaDB)
	DataNascimento string `json:"datanascimento" ` //maximo 8 digitos
	Name           string `json:"name" binding:"required"`
	Lastname       string `json:"lastname" binding:"required"`
}

func (p *Profile) Validate(action string) error {
	switch strings.ToLower(action) {
	case "createprofile":
		if len(p.Nickname) > 15 {
			return errors.New("Nickname Invalido maior que 15")
		}

	case "login":

	default:

	}
	return nil
}
