<<<<<<< Updated upstream
package profile

type Profile struct{
	
		id int32;
 		avatar string //longtext
	 	telefone int8 //a dividir
		data_nascimento int8 //maximo 8 digitos
		idade 
	
} 
=======
package models

type Profile struct {
	ID       uint32  `json:"id" gorm:"type:bigint;primary_key; AUTO_INCREMENT"`
	Avatar         string  `json:"avatar" gorm:"type:longtext"` //longtext no BD (mysql-MariaDB)
	DataNascimento string  `json:"datanascimento" `//maximo 8 digitos
	Nickname       string  `json:"nickname" `
}
>>>>>>> Stashed changes
