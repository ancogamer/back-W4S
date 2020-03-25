package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claim is the token payload
type Claim struct {
	User User `json:"user,omitempty"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	jwt.StandardClaims
}
