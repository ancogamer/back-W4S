package models

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	UserEmail string `json:"user,omitempty"`
	jwt.StandardClaims
}
