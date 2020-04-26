package models

type TOTPkey struct {
	ID        uint64 `json:"id" gorm:"type:bigint;primary_key; AUTO_INCREMENT"`
	Key       string `json:"totpkey"`
	UserEmail string `json :"email"`
	Actived   int64  `json:"actived"`
}
