package models

import "github.com/jinzhu/gorm"

type Table struct {
	gorm.Model
	Name                 string     `json:"name"`
	Description          string     `json:"description"`
	NumberOfParticipants int        `json:"numberofparticipants"`
	MaxOfParticipants    int        `json:"maxofparticipants"`
	Thumbnail            string     `json:"thumbnail"`
	AdventureLink        string     `json:"adventurelink" gorm:"foreignkey:TableID"`
	RpgSystem            RpgSystem  `json:"rpgsystem" gorm:"foreignkey:TableID"`
	OthersLinks          OtherLinks `json:"otherlinks" gorm:"foreignkey:TableID"`
	Pictures             []*Picture `json:"pictures" gorm:"many2many:user_languages;foreignkey:TableID"`
	Posts                []*Post    `json:"posts" gorm:"many2many:user_languages;foreignkey:TableID"`
	User                 []*User    `json:"idusers" gorm:"many2many:user_Tables;ForeignKey:id;AssociationForeignKey:id"`
}

type TableInput struct {
	Name                 string     `json:"name" binding:"required"`
	Description          string     `json:"description" binding:"required"`
	NumberOfParticipants int        `json:"numberofparticipants"`
	MaxOfParticipants    int        `json:"maxofparticipants"`
	Thumbnail            string     `json:"thumbnail"`
	AdventureLink        string     `json:"adventurelink" gorm:"foreignkey:TableID"`
	RpgSystem            RpgSystem  `json:"rpgsystem" gorm:"foreignkey:TableID"`
	OthersLinks          OtherLinks `json:"otherlinks" gorm:"foreignkey:TableID"`
	Pictures             []Picture  `json:"pictures" gorm:"foreignkey:TableID"`
	Posts                []Post     `json:"posts" gorm:"foreignkey:TableID"`
	IDUsers              []uint     `json:"idusers" gorm:"foreignkey:ID"`
}
