//Authentication of package
package authc

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
	"w4s/models"
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
	fmt.Println("gerando token")
	user.Token=""
	claims:=models.Claim{
		UserEmail:user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		StandardClaims: jwt.StandardClaims{
			Issuer:"system",
			ExpiresAt:time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenpass, err:=os.LookupEnv("TOKEN_PASSWORD")
	fmt.Println(tokenpass)
	if err {
		fmt.Println("sdasdas",tokenpass)
	}
	return token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
}

// ValidateToken validate a JWT

func ValidateToken(c *gin.Context) bool {
	userToken := c.Request.Header.Get("Authorization")
	detail := TokenDetail{
		Valid:  true,
	//	Active: true,
	}
	split := strings.Split(userToken, " ")
	if len(split) != 2 || split[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não é um token Bearer",
		})
		return false
	}
	token, err := jwt.ParseWithClaims(split[1], &models.Claim{}, func(token *jwt.Token) (interface{}, error) { return os.Getenv("TOKEN_PASSWORD"), nil })
	/*if err != nil || token.Valid != true {
		c.JSON(http.StatusOK, gin.H{
			"error":"token invalidou ou expirado",
		})
		return false
	}*/
	if token.Valid {
		fmt.Println("You look nice today")
		return true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		detail.Valid = false
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			c.JSON(http.StatusUnauthorized, errors.New("Token expirado"))
			//detail.Active = false
			return false

		} else {
			fmt.Println("Couldn't handle this token:", err)
			c.JSON(http.StatusUnauthorized, err)
			c.JSON(http.StatusUnauthorized, gin.H{"Couldn't handle this token": err})
			return false
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"Couldn't handle this token": err})
		detail.Valid = false
		//detail.Active = false
		return false
	}
	return detail.Valid
}
