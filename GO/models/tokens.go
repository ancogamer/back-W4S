package models

import "github.com/jinzhu/gorm"

type LogoffListTokens struct {
	gorm.Model
	Token string `json:"token"`
}
type AccountCreatedToken struct {
	gorm.Model
	Token string `json:"token"`
}
