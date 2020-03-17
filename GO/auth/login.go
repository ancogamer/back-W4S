package auth

import (
	"github.com/jinzhu/gorm"
	"github.com/jw"
	"w4s/models"
)
//login 
// SignIn method
func SignIn(email, password string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		user.Password = ""
		return GenerateJWT(user)
	}

	return "", err
}
