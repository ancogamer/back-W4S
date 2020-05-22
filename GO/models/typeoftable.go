package models

import "github.com/jinzhu/gorm"

type TypeofTable struct {
	gorm.Model
	TableID     uint `json:"id_table"`
	Name        string
	Description string
}
