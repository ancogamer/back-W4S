package models

import "github.com/jinzhu/gorm"

type PermissionTable struct {
	gorm.Model
	Permission      int  `json:"permission"`
	ProfileNickname uint `json:"-"`
	TableId         uint `json:"-"`
}
