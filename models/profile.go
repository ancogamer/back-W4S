package models

import (
	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	IDUser         uint    `json:"id_user" form:"id_user" gorm:"default:0"`
	Nickname       string  `json:"nickname" form:"nickname" gorm:"unique_index"` //max 15
	Name           string  `json:"name" form:"name"`
	Lastname       string  `json:"lastname" form:"lastname"`
	Avatar         string  `json:"avatar" form:"avatar" gorm:"type:text"` //longtext no BD (mysql-MariaDB)
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
