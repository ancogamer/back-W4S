package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/smtp"
	"os"
	"w4s/authc"
)

func ConfirmationEmail(userEmail string, c *gin.Context) {
	passcode := authc.TOTPGenerate(userEmail, c)
	//Thanks https://blog.mailtrap.io/golang-send-email/#Sending_emails_with_smtpSendMail for the code
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "findatablew4s@gmail.com", os.Getenv("EMAIL_PASSWORD"), "smtp.gmail.com")
	// Here we do it all: connect to our server, set up a message and send it
	to := []string{userEmail}
	msg := []byte("To: " + userEmail + "\r\n" +
		"Subject: NÂO RESPONDA ESTE EMAIL - Confirmação de Conta \r\n" +
		"\r\n" +
		"Bem vindo ! obrigado por criar uma conta no Find a Table - RPG !\r\n" +
		"\r\n" +
		"Seu Confirme o codigo no seu aplicativo " + passcode +
		"\r\n" +
		"Caso não tenha criado, por favor entrar em contato com findatablew4s@gmail.com, com o Assunto : Conta Criada Indevidamente ")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "findatablew4s@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
