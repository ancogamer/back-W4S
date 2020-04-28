package controllers

import (
	"github.com/gin-gonic/gin"
	"net/smtp"
	"os"
	"w4s/authc"
)

func ConfirmationEmail(userEmail string, c *gin.Context) {
	authc.GenerateJWT(userEmail,14400)
	SendEmail(userEmail)
/*	if e := config.SendMail([]string{user.EmailID}, "Verification", emailBody); e != nil {
		err = fmt.Errorf("Error in sending mail %v", e)
		log.Println(err)
		// since there was error in sending mail
		// hence the user can't sign up
		// so need to delete the user from db
		// so that he/she can try again
		// delete the user details from the db
		if err := s.DB.DeleteUserDetails(user.EmailID); err != nil {
			log.Println("Error in deleting user details upon failed verification: ", err)
		}
		http.Error(w, err.Error(), 500)*/
		return

}
func SendEmail(userEmail string, userSingup string) error{
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
		"Clique aqui para confirmar sua conta = " + userSingup +
		"\r\n" +
		"Caso não consiga, é só copiar o link e colar no navegador ! "+
		"\r\n"+
		"Caso não tenha criado, por favor entrar em contato com findatablew4s@gmail.com, com o Assunto : Conta Criada Indevidamente ")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "findatablew4s@gmail.com", to, msg)
	if err != nil {
		return err
	}
	return nil
}