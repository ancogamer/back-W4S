//Database connections with gorm
package DB

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"w4s/models"
)

//	_ "github.com/jinzhu/gorm/dialects/mysql"
func SetupModels() *gorm.DB {

	/*
		db, err := gorm.Open("mysql",
		""+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@/w4s?charset=utf8&parseTime=True&loc=Local")
	*/

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.UserAccountBadListToken{})
	db.AutoMigrate(&models.Table{})
	db.AutoMigrate(&models.Picture{})
	db.AutoMigrate(&models.OtherLinks{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Profile{})
	db.AutoMigrate(&models.TypeofTable{})
	db.AutoMigrate(&models.RpgSystem{})
	db.AutoMigrate(&models.LogoffListTokens{})
	db.AutoMigrate(&models.User{})

	return db
}
