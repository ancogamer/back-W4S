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
func GenerateJWT(user models.User, experatingtime time.Duration) (string, error) {
	// Create the Claims
	claims := models.Claim{
		User: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * experatingtime).Unix(),
			Issuer:    "Find A Table System",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
}

// ValidateToken validate a JWT
func ValidateToken(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tokenLogOffList models.LogoffListTokens
	userToken := c.Request.Header.Get("Authorization")
	if userToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "não logado",
		})
		return
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
			return
		}
		//Verifing the token/ Verificando o token
		token, err := jwt.ParseWithClaims(split[1], &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
		})
		_, ok := token.Claims.(*models.Claim)
		if ok && token.Valid {
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": "Não foi possível autenticar", "error ": err})
		return
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "não logado",
	})

	/*//IF YOU WANT RETURN ERROS OH A SEPARATED WAY// Caso você queira retornar erros de maneira separada
	if token.Valid {
		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"error ":"NOT A TOKEN"})
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"error ":"TOKEN EXPIRED"})
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"error ":"Couldn't handle this token:"})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error ":"Couldn't handle this token:",
			"err":err,
		})
	}*/
	return
}
