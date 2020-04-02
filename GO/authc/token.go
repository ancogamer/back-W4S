//Authentication of package
package authc

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
	return token.SignedString([]byte(os.Getenv("token_password")))
}

// ValidateToken validate a JWT
func ValidateToken(user models.User,c *gin.Context) (bool, error) {
	valid, err := TokenInfos(user.Token, c)
	if err != nil {
		return false, err
	}
	ret := valid.Active
	return ret, nil
}

// TokenInfos return ifos of JWT
func TokenInfos(tokenString string, c *gin.Context) (TokenDetail, error) {
	// mySigningKey := []byte(secret)
	// Token from another example.  This token is expired
	fmt.Println(tokenString)
	detail := TokenDetail{
		Valid:  true,
		Active: true,
	}
	//Review this part ! -REVER ESTA PARTE
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("fiscaluno"), nil
	})
	if token.Valid {
		fmt.Println("Acesso permitido")
		detail.Note = "1"
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		detail.Valid = false
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("Token invalido")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Token de acesso invalido",
			})
			detail.Valid = false
			detail.Active = false
			detail.Note = "2"
			return detail, nil
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Token expirado")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Seu acesso expirou, por favor entre novamente",
			})
			detail.Active = false
			detail.Note = "3"
			return detail, nil
		} else {
			fmt.Println("1-Couldn't handle this token:", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			detail.Active = false
			detail.Note = err.Error()
			return detail, err
		}
	} else {
		fmt.Println("2-Couldn't handle this token:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		detail.Valid = false
		detail.Active = false
		detail.Note = err.Error()
		return detail, err
	}
	return detail, nil
}
