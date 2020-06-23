package models

import "github.com/jinzhu/gorm"

type Table struct {
	gorm.Model
	Thumbnail            string            `json:"thumbnail"`                                                                             //capa
	Name                 string            `json:"name"`                                                                                  //nome
	Description          string            `json:"description"`                                                                           //descrição
	NumberOfParticipants int               `json:"numberofparticipants"`                                                                  //numero de participantes
	MaxOfParticipants    int               `json:"maxofparticipants"`                                                                     //numero maximo de participantes
	Privacy              int               `json:"privacy" gorm:"default:1"`                                                              //privacidade
	RpgSystem            string            `json:"rpgsystem" gorm:"foreignkey:TableID"`                                                   //sistema
	Links                string            `json:"otherlinks" gorm:"foreignkey:TableID"`                                                  //links 	//images
	User                 []*Profile        `json:"idusers" gorm:"default:0;many2many:user_Tables;ForeignKey:id;AssociationForeignKey:id"` // Perfil
	Permitions           []PermissionTable `json:"permissions"`
	//Posts                []*Post    `json:"posts" gorm:"many2many:user_languages;foreignkey:TableID"`
	//AdventureLink        string     `json:"adventurelink" gorm:"foreignkey:TableID"`
}

type TableInput struct {
	Name                 string            `json:"name" binding:"required"`
	Description          string            `json:"description" binding:"required"`
	NumberOfParticipants int               `json:"numberofparticipants"`
	MaxOfParticipants    int               `json:"maxofparticipants"`
	Thumbnail            string            `json:"thumbnail"`
	Privacy              int               `json:"privacy" gorm:"default:1"`                                                    //privacidade
	RpgSystem            string            `json:"rpgsystem" gorm:"foreignkey:TableID"`                                         //sistema
	Links                string            `json:"otherlinks" gorm:"foreignkey:TableID"`                                        //links
	User                 []*Profile        `json:"idusers" gorm:"many2many:user_Tables;ForeignKey:id;AssociationForeignKey:id"` // Perfil
	Permitions           []PermissionTable `json:"permissions"`
	//Posts                []*Post    `json:"posts" gorm:"many2many:user_languages;foreignkey:TableID"`
	//AdventureLink        string     `json:"adventurelink" gorm:"foreignkey:TableID"`
}
