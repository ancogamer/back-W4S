package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/smtp"
	"os"
	"w4s/authc"
	"w4s/models"
)

/*func SendConfirmationChangeEmail(userEmail string, c*gin.Context) error{
	db:=c.MustGet("db").(*gorm.DB)
	var user models.User
	if err:=db.Where("email = ?",userEmail).Find(&user);err!=nil{
		userSingUpToken, err := authc.GenerateJWT(userEmail, 14400)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return err
		}
		userURL := "http://localhost:8080/user/confirm_email_change?e=" + userEmail + "&t=" + userSingUpToken
		//Thanks https://blog.mailtrap.io/golang-send-email/#Sending_emails_with_smtpSendMail for the code
		// Choose auth method and set it up
		msg := []byte("To: " + userEmail + "\r\n" +
			"Subject: NÂO RESPONDA ESTE EMAIL - Confirmação de Troca de Email \r\n" +
			"\r\n" +
			"Clique aqui para confirmar a troca do email = " + userURL +
			"\r\n" +
			"Caso não consiga, é só copiar o link e colar no navegador ! " +
			"\r\n" +
			"Caso não tenha pedido esta troca, por favor sugerimos mudar a sua senha e  entrar em contato com findatablew4s@gmail.com, com o Assunto : Troca de email indevida")
		if err := SendEmail(userEmail,msg); err != nil {
			//Declaring and inicializing a userAccountCreatedToken
			//Declarando e inicializando uma variavel userAccount
			userAccountCreatedToken := models.AccountCreatedToken{
				Token: userSingUpToken,
			}
			db := c.MustGet("db").(*gorm.DB) //Establish conection with database/Estabelecendo conexão com o banco de dados
			if err := db.Create(userAccountCreatedToken).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
				return err
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return err
		}
		return nil
	}
}
*/
func SendConfirmationCreateAccountEmail(userEmail string, c *gin.Context) error {
	userSingUpToken, err := authc.GenerateJWT(userEmail, 14400)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return err
	}
	userURL := "http://localhost:8080/user/confirm?e=" + userEmail + "&t=" + userSingUpToken
	//Thanks https://blog.mailtrap.io/golang-send-email/#Sending_emails_with_smtpSendMail for the code
	// Choose auth method and set it up
	msg := []byte("To: " + userEmail + "\r\n" +
		"Subject: NÂO RESPONDA ESTE EMAIL - Confirmação de Conta \r\n" +
		"\r\n" +
		"Bem vindo ! obrigado por criar uma conta no Find a Table - RPG !\r\n" +
		"\r\n" +
		"Clique aqui para confirmar sua conta = " + userURL +
		"\r\n" +
		"Caso não consiga, é só copiar o link e colar no navegador ! " +
		"\r\n" +
		"Caso não tenha criado, por favor entrar em contato com findatablew4s@gmail.com, com o Assunto : Conta Criada Indevidamente ")
	if err := SendEmail(userEmail, msg); err != nil {
		//Declaring a new user to be populated/
		//Declarando um novo usuario que será populado
		var user models.User
		//Declaring and inicializing a userAccountCreatedToken
		//Declarando e inicializando uma variavel userAccount
		userAccountCreatedToken := models.AccountCreatedToken{
			Token: userSingUpToken,
		}
		db := c.MustGet("db").(*gorm.DB) //Establish conection with database/Estabelecendo conexão com o banco de dados
		if err := db.Create(userAccountCreatedToken).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return err
		}
		if err := db.Where("email = ?", userEmail).Find(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return err
		}
		fmt.Println(user)
		db.Unscoped().Delete(&user)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return err
	}
	return nil
}
func SendEmail(userEmail string, msg []byte) error {
	auth := smtp.PlainAuth("", "findatablew4s@gmail.com", os.Getenv("EMAIL_PASSWORD"), "smtp.gmail.com")
	// Here we do it all: connect to our server, set up a message and send it
	to := []string{userEmail}
	err := smtp.SendMail("smtp.gmail.com:587", auth, "findatablew4s@gmail.com", to, msg)
	if err != nil {
		return err
	}
	return nil
}
