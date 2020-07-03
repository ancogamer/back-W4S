package models

import "github.com/jinzhu/gorm"

type PermissionTable struct {
	gorm.Model
	Permission      string `json:"permission"`
	ProfileNickname uint   `json:"nickname"`
	TableId         uint   `json:"-"`
}
