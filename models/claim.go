package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	UserEmail string `json:"user,omitempty"`
	jwt.StandardClaims
}

func (c *Claim) Validate() error {
	if err := checkmail.ValidateFormat(c.UserEmail); err != nil {
		return errors.New("Digite um endereço de e-mail válido")
	}
	return nil
}
