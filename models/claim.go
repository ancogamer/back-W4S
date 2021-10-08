package models

import (
	"errors"

	"github.com/badoux/checkmail"
	"github.com/golang-jwt/jwt"
)

type Claim struct {
	UserEmail string `json:"user,omitempty"`
	jwt.StandardClaims
}

func (c *Claim) Validate() error {
	if err := checkmail.ValidateFormat(c.UserEmail); err != nil {
		return errors.New("digite um endereço de e-mail válido")
	}
	return nil
}
