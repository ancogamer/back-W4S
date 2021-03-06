//Authentication of package
package authc

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"strings"
	"time"
	"w4s/models"
)

// GenerateJWT creates a new token to the client
func GenerateJWT(userEmail string, experatingTime time.Duration) (string, error) {
	// Create the Claims
	claims := models.Claim{
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * experatingTime).Unix(),
			Issuer:    "Find A Table System",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
}

// ValidateToken validate a JWT
func ValidateLoginToken(c *gin.Context) string {
	db := c.MustGet("db").(*gorm.DB)
	var tokenLogOffList models.LogoffListTokens
	userToken := c.Request.Header.Get("Authorization")
	//userEmail:=c.Query()
	if userToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "não logado",
		})
		return ""
	}
	//Checking if the token are on the Logoff list/ Checando se o token esta na lista de deslogados
	if err := db.Where("token = ?", userToken).Find(&tokenLogOffList).Error; gorm.IsRecordNotFoundError(err) {
		//If this Record Was not found, it means that the user is loged/ Se o registro não foi encontrado, significa que o usuario esta logado
		//Spliting the token/ Separando o token
		split := strings.Split(userToken, " ")
		if len(split) != 2 || split[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "não é um token Bearer",
			})
			return ""
		}
		//Verifing the token/ Verificando o token
		claim, err := ValidateToken(split[01])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return ""
		}
		return claim
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "não logado",
	})
	return ""
}

//This Parse the token and check if is valid :D
func ValidateToken(userToken string) (string, error) {
	token, err := jwt.ParseWithClaims(userToken, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_PASSWORD")), nil
	})
	if claims, ok := token.Claims.(*models.Claim); ok && token.Valid {
		return claims.UserEmail, nil
	}
	return "", err
}
