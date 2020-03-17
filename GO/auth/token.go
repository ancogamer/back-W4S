package auth

import (
	"errors"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/alfredomendozacap/jwt-auth/config"
	"github.com/alfredomendozacap/jwt-auth/api/models"
	"github.com/alfredomendozacap/jwt-auth/api/responses"
	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT creates a new token to the client
func GenerateJWT(user models.User) (string, error) {
	claim := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Alfredo Mendoza",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(config.SECRETKEY)
}

// ExtractToken retrieves the token from headers ans Query
func ExtractToken(w http.ResponseWriter, r *http.Request) *jwt.Token {
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return config.SECRETKEY, nil
		},
	)

	if err != nil {
		code := http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				err = errors.New("Your token has expired")
				responses.ERROR(w, code, err)
				return nil
			case jwt.ValidationErrorSignatureInvalid:
				err = errors.New("The signature is invalid")
				responses.ERROR(w, code, err)
				return nil
			default:
				responses.ERROR(w, code, err)
				return nil
			}
		}
	}

	return token
}
