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
func ValidateLoginToken(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tokenLogOffList models.LogoffListTokens
	userToken := c.Request.Header.Get("Authorization")
	//userEmail:=c.Query()
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
		//This was seperated, to be possible use this on other places/
		//Isto foi separado para poder ser utilizado em outros lugares
		err := ValidateToken(split[01])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		}
		return
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "não logado",
	})
	return
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
}

//This Parse the token and check if is valid :D
func ValidateToken(userToken string) error {
	token, err := jwt.ParseWithClaims(userToken, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_PASSWORD")), nil
	})
	_, ok := token.Claims.(*models.Claim)
	if ok && token.Valid {
		return nil
	}
	return err
}
