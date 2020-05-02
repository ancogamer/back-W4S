package models

import "github.com/jinzhu/gorm"

type LogoffListTokens struct {
	gorm.Model
	Token string `json:"token"`
}
type UserAccountBadListToken struct {
	gorm.Model
	Token string `json:"token"`
}
