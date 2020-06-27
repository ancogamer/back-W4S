package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type Profile struct {
	gorm.Model
	IDUser         uint    `json:"id_user" form:"id_user" gorm:"default:0"`
	Nickname       string  `json:"nickname" form:"nickname" gorm:"unique_index"` //max 15
	Name           string  `json:"name" form:"name"`
	Lastname       string  `json:"lastname" form:"lastname"`
	Avatar         string  `json:"avatar" form:"avatar"`                  //longtext no BD (mysql-MariaDB)
	DataNascimento string  `json:"datanascimento" form:"datanascimento" ` //maximo 8 digitos
	Deleted        bool    `json:"deleted" form:"deleted" gorm:"type:BOOLEAN"`
	Tables         []Table `json:"tables" gorm:"many2many:user_Tables;ForeignKey:id;AssociationForeignKey:id"`
}

type ProfileInput struct {
	Nickname       string `json:"nickname" form:"nickname" binding:"required" `
	Avatar         string `json:"avatar" form:"avatar"`                  //longtext no BD (mysql-MariaDB)
	DataNascimento string `json:"datanascimento" form:"datanascimento" ` //maximo 8 digitos
	Name           string `json:"name" form:"name" binding:"required"`
	Lastname       string `json:"lastname" form:"lastname" binding:"required"`
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
