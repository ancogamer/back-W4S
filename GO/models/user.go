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
	Nickname string  `json:"nickname "`
	Email    string  `json:"email" gorm:"type:varchar(100);unique_index" `
	Password string  `json:"-"`
	Name     string  `json:"name"`
	Lastname string  `json:"lastname"`
	Deleted  bool    `json:"deleted" gorm:"type:BOOLEAN"`
	Actived  bool    `json:"actived" gorm:"type:BOOLEAN"`
	Profile  Profile `json:"profile,omitempty" gorm:"foreignkey:IdUser"`
	Token    string  `json:"token";sql:"-"`
}

//With biding required in all fields/ Com o biding obrigatorio em todos os campos
type UserInput struct {
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required" `
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
}
type UserInputRecoveryPassword struct {
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
	Name            string `json:"name"`
	Lastname        string `json:"lastname"`
}

//With out the biding required in all fields/ Sem o biding obrigatorio em todos os campos
type UserInputUpdate struct {
	Nickname           string `json:"nickname"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	NewPassword        string `json:"newpassword"`
	ConfirmNewPassword string `json:"confirmnewpassword"`
	Name               string `json:"name"`
	Lastname           string `json:"lastname"`
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
	case "updateEmail":
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Digite um endereço de e-mail válido")
		}
	case "login":
		if u.Nickname == "" && u.Email == "" {
			return errors.New("Preencha este campo")
		}
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
