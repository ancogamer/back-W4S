package models

import "github.com/jinzhu/gorm"

type RpgSystem struct {
	gorm.Model
	TableID     uint   `json:"id_table"`
	Name        string `json:"rpgsystemname"`
	Description string `json:"rpgsystemdescripton"`
}
