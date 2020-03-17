package models

type Profile struct {
	ID             uint32  `json:"id" gorm:"type:bigint;primary_key; AUTO_INCREMENT"`
	Avatar         string  `json:"avatar" gorm:"type:longtext"` //longtext no BD (mysql-MariaDB)
	DataNascimento string  `json:"datanascimento" `//maximo 8 digitos
	Nickname       string  `json:"nickname" `
	UserID         uint32  `gorm:"not null" json:"author_id"`
	User 		   User    `gorm:"foreignkey:UserID" json:"user"`
}

/*
	Author    User      `gorm:"foreignkey:AuthorID" json:"author"`
	AuthorID  uint64    `gorm:"not null" json:"author_id"`
*/