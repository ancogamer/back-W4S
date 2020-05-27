package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
	"unicode"
	"w4s/security"

	"github.com/badoux/checkmail"
)

//Struct from the User
type User struct {
	gorm.Model
	Email     string  `json:"email" gorm:"type:varchar(100);unique_index" `
	Password  string  `json:"-"`
	Deleted   bool    `json:"deleted" gorm:"type:BOOLEAN"`
	Actived   bool    `json:"actived" gorm:"type:BOOLEAN"`
	ProfileID uint    `json:"profileid"`
	Profile   Profile `json:"profile"`
	Tables    []Table `json:"tables" gorm:"many2many:user_Tables;ForeignKey:id;AssociationForeignKey:id"`
	Token     string  `json:"token"`
}

//With biding required in all fields/ Com o biding obrigatorio em todos os campos
type UserInput struct {
	Email    string `json:"email" binding:"required" `
	Password string `json:"password" binding:"required"`
}

//Separete User Input to recovery password
type UserInputRecoveryPassword struct {
	Email           string `json:"email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
}

//With out the biding required in all fields/ Sem o biding obrigatorio em todos os campos
type UserInputUpdate struct {
	Email              string `json:"email"`
	Password           string `json:"password"`
	NewPassword        string `json:"newpassword"`
	ConfirmNewPassword string `json:"confirmnewpassword"`
}

// BeforeSave hash the user password
func BeforeSave(password string) (string, error) {
	hashedPassword, err := security.Hash(password)
	if err != nil {
		panic("Password hash")
	}
	password = string(hashedPassword)
	return password, nil
}

// Validate validates the inputs
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "createuser":
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Digite um endereço de e-mail válido")
		}
		if err := PasswordCheck(u.Password); err != nil {
			return errors.New(err.Error())
		}
	case "updateemailandresendlink":
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Digite um endereço de e-mail válido")
		}
	case "login":
		if err := checkmail.ValidateFormat(u.Email); u.Email != "" && err != nil {
			return errors.New("Invalid email")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
	default:
	}
	return nil
}

func PasswordCheck(password string) error {
	if password == "" {
		return errors.New("Password is required")
	}
	if len(password) > 20 {
		return errors.New("Insira uma senha valida")
	}
	if err := ValidatorPassword(password); err != true {
		return errors.New("Insira uma senha valida")
	}
	return nil
}
func ValidatorPassword(pass string) bool {
	//Thanks to http://www.inanzzz.com/index.php/post/8l1a/validating-user-password-in-golang-requests
	//Password validates plain password against the rules defined below.
	//
	// upp: at least one upper case letter.
	// low: at least one lower case letter.
	// num: at least one digit.
	// sym: at least one special character.
	// tot: at least eight characters long.
	// No empty string or whitespace.
	var (
		upp, low, num, sym bool
		tot                uint8
	)
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}
	if !upp || !low || !num || !sym || tot < 5 {
		return false
	}
	return true
}
