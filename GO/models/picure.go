package models

import "github.com/jinzhu/gorm"

type Picture struct {
	gorm.Model
	TableID     uint   `json:"id_table"`
	PictureFile string `json:"picturefile"` //Picture on base 64
}
