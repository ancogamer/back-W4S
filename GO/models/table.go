package models

import "github.com/jinzhu/gorm"

type Mesa struct {
	gorm.Model
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Thumbnail     string     `json:"thumbnail"`
	AdventureLink string     `json:"adventurelink" gorm:"foreignkey:TableID"`
	RpgSystem     RpgSystem  `json:"rpgsystem" gorm:"foreignkey:TableID"`
	OthersLinks   OtherLinks `json:"otherlinks" gorm:"foreignkey:TableID"`
	Pictures      []Picture  `json:"pictures" gorm:"foreignkey:TableID"`
	Posts         []Post     `json:"posts" gorm:"foreignkey:TableID"`
}
