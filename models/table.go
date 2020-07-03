package models

import "github.com/jinzhu/gorm"

type Table struct {
	gorm.Model
	Thumbnail            string            `json:"thumbnail" form:"thumbnail"`                                                                          //capa
	Name                 string            `json:"name" form:"name" gorm:"unique_index"`                                                                //nome
	Description          string            `json:"description" form:"description" `                                                                     //descrição
	NumberOfParticipants int               `json:"numberofparticipants" form:"numberofparticipants"`                                                    //numero de participantes
	MaxOfParticipants    int               `json:"maxofparticipants" form:"maxofparticipants"`                                                          //numero maximo de participantes
	Privacy              int               `json:"privacy" form:"privacy" gorm:"default:1"`                                                             //privacidade
	RpgSystem            string            `json:"rpgsystem" form:"rpgsystem" gorm:"foreignkey:TableID"`                                                //sistema
	Links                string            `json:"otherlinks" form:"otherlinks" gorm:"foreignkey:TableID"`                                              //links 	//images
	User                 []*Profile        `json:"idusers" form:"iduser" gorm:"default:0;many2many:user_Tables;ForeignKey:id;AssociationForeignKey:id"` // Perfil
	Permitions           []PermissionTable `json:"permissions" form:"permissions"`
	//Posts                []*Post    `json:"posts" gorm:"many2many:user_languages;foreignkey:TableID"`
	//AdventureLink        string     `json:"adventurelink" gorm:"foreignkey:TableID"`
}

type TableInput struct {
	Name                 string            `json:"name" form:"name" binding:"required"`
	Description          string            `json:"description" form:"description" binding:"required"`
	NumberOfParticipants int               `json:"numberofparticipants" form:"numberofparticipants"`
	MaxOfParticipants    int               `json:"maxofparticipants" form:"maxofparticipants"`
	Thumbnail            string            `json:"thumbnail" form:"thumbnail"`
	Privacy              int               `json:"privacy" form:"privacy" gorm:"default:1"`                                                    //privacidade
	RpgSystem            string            `json:"rpgsystem" form:"rpgsystem" gorm:"foreignkey:TableID"`                                       //sistema
	Links                string            `json:"otherlinks" form:"otherlinks" gorm:"foreignkey:TableID"`                                     //links
	User                 []*Profile        `json:"idusers" form:"idusers" gorm:"many2many:user_Tables;ForeignKey:id;AssociationForeignKey:id"` // Perfil
	Permitions           []PermissionTable `json:"permissions" form:"permissions"`
	//Posts                []*Post    `json:"posts" gorm:"many2many:user_languages;foreignkey:TableID"`
	//AdventureLink        string     `json:"adventurelink" gorm:"foreignkey:TableID"`
}
