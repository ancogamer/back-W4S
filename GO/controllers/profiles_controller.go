package controllers
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)

type CreateProfileInput struct{
	Avatar         string  `json:"avatar" binding:required` //longtext no BD (mysql-MariaDB)
	DataNascimento string  `json:"datanascimento" binding:required`//maximo 8 digitos
	Nickname       string  `json:"nickname" binding:required `
}
func CreateProfile( c *gin.Context){
	db:= c.MustGet("db").(*gorm.DB)
	//Validating input
	var input CreateProfileInput

	if err:= c.ShouldBindJSON(&input);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		defer db.Close()
		return
	}
	//Creating Profile
	profile:=models.Profile{
		Avatar:         input.Avatar,
		DataNascimento: input.DataNascimento,
		Nickname:       input.Nickname,
	}
	//Validation part.
	/*err := profile.Validate("")
	if err!=nil{
		defer db.Close()
		c.JSON(http.StatusNotImplemented, gin.H{
			"err":err,
		})
	}else{*/
			if dbc := db.Create(&profile); dbc.Error != nil { //return the error by JSON
				defer db.Close()
				c.JSON(http.StatusBadRequest, gin.H{"erro": dbc.Error})
			} else { //return the post data if is ok, by JSON
				defer db.Close()
				c.JSON(http.StatusOK, gin.H{"data": profile})
			}
		}
	//}
}
//Find all users on the database
func FindProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var profile []models.Profile
	db.Find(&profile)
	defer db.Close()
	if profile!=nil{//checking if something was returned.
		c.JSON(http.StatusNotFound, gin.H{
			"error":"nenhum registro encontrado",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"user":profile,
		})
	}
}
