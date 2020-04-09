//Authentication of package
package authc

import (
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
/*type Token struct {
	Raw       string                 // The raw token.  Populated when you Parse a token
	Header    map[string]interface{} // The first segment of the token
	Claims    models.Claim                 // The second segment of the token
	Signature string                 // The third segment of the token.  Populated when you Parse a token
	Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
}
*/

// GenerateJWT creates a new token to the client
func GenerateJWT(user models.User) (string, error) {
	fmt.Println("gerando token")
	user.Token=""
	/*claims:=models.Claim{
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

	 */

	// Create the Claims
	claims := models.Claim{
		UserEmail:      "",
		CreatedAt:      time.Time{},
		UpdatedAt:      time.Time{},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "system",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

}

// ValidateToken validate a JWT

func ValidateToken(c *gin.Context) bool {
	userToken := c.Request.Header.Get("Authorization")
	if userToken ==""  {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Usuario":"não logado",
		})
		return false
	}

	split := strings.Split(userToken, " ")
	if len(split) != 2 || split[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "não é um token Bearer",
		})
		return false
	}

	token, err := jwt.ParseWithClaims(userToken, &models.Claim{}, func(token *jwt.Token) (interface{}, error) { return os.Getenv("TOKEN_PASSWORD"), nil })
	if _, ok := token.Claims.(*models.Claim); ok && token.Valid {
		return true
	} else {
		c.JSON(http.StatusBadRequest,err)
		return  false
	}
}
