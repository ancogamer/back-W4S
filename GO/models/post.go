package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	TableID  uint   `json:"id_table"`
	PostInfo string `json:"postinfo"`
}
