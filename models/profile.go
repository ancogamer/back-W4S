package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Profile struct {
	gorm.Model
	IDUser         uint      `json:"id_user"`
	Avatar         string    `json:"avatar"`          //longtext no BD (mysql-MariaDB)
	DataNascimento time.Time `json:"datanascimento" ` //maximo 8 digitos
}
type ProfileInput struct {
	Avatar         string    `json:"avatar"`          //longtext no BD (mysql-MariaDB)
	DataNascimento time.Time `json:"datanascimento" ` //maximo 8 digitos
}

/*
Case need
func (p *Profile) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":

	case "login":

	default:

	}
	return nil
}
*/
