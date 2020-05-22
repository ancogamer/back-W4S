package models

import "github.com/jinzhu/gorm"

type OtherLinks struct {
	gorm.Model
	TableID uint        `json:"id_table"`
	Discord string      `json:"discord"`
	Skype   string      `json:"skype"`
	Others  []OtherLink `json:"others" gorm:"foreignkey:OtherLinkID"`
}
type OtherLink struct {
	OtherLinkID uint
	Link        string
}
