package models

type Profile struct {
	ID             uint32  `json:"id" gorm:"type:bigint;primary_key; AUTO_INCREMENT"`
	Avatar         string  `json:"avatar" gorm:"type:longtext"` //longtext no BD (mysql-MariaDB)
	DataNascimento string  `json:"datanascimento" `//maximo 8 digitos
	Nickname       string  `json:"nickname" `
	User           uint32  `json:"id_users" gorm:"type:bigint;foreign key(id_users) references users(id)"`
}

