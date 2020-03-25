//Authentication of package
package authc

import (
	"fmt"
	"os"
	"time"
	"w4s/models"
	jwt "github.com/dgrijalva/jwt-go"
)
//Detais from the token struct
type TokenDetail struct {
	Valid  bool   `json:"valid"`
	UserID uint   `json:"userID,omitempty"`
	Active bool   `json:"active"`
	Note   string `json:"note,omitempty"`
}

// GenerateJWT creates a new token to the client
func GenerateJWT(user models.User) (string, error) {
	user.Token=""
	claims:=models.Claim{
		User:user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		StandardClaims: jwt.StandardClaims{
			Issuer:"system",
			ExpiresAt:time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("token_password")))
}
// ValidateToken validate a JWT
func ValidateToken(user models.User, secret string) (bool, error) {
	valid, err := TokenInfos(user.Token)
	if err != nil {
		return false, err
	}
	ret := valid.Active
	return ret, nil
}

// TokenInfos return ifos of JWT
func TokenInfos(tokenString string) (TokenDetail, error) {
	// mySigningKey := []byte(secret)
	// Token from another example.  This token is expired

	detail := TokenDetail{
		Valid:  true,
		Active: true,
	}
	//Review this part ! -REVER ESTA PARTE
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("fiscaluno"), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
		detail.Note = "You look nice today"
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		detail.Valid = false
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			detail.Valid = false
			detail.Active = false
			detail.Note = "That's not even a token"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			detail.Active = false
			detail.Note = "Timing is everything"
			return detail, nil

		} else {
			fmt.Println("Couldn't handle this token:", err)
			detail.Active = false
			detail.Note = err.Error()
			return detail, err

		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		detail.Valid = false
		detail.Active = false
		detail.Note = err.Error()
		return detail, err
	}

	return detail, nil
}
