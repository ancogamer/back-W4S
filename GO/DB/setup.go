//Database connections with gorm
package DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"w4s/models"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql",
		"Saletti:Saletti123@/w4s?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Profile{})
	return db
}
