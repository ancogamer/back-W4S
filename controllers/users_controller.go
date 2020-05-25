//POST /user
//Create a new user
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/authc"
	"w4s/handlers"
	"w4s/models"
	"w4s/security"
)

//Create User
func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//Validating input
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	if db.Where("nickname = ?", input.Nickname).Find(&user).RecordNotFound() {
		//Creating user
		user.Nickname = input.Nickname
		user.Email = input.Email
		user.Password = input.Password
		user.Name = input.Name
		user.Lastname = input.Lastname
		user.Actived = false
		user.Deleted = false
		user.Token = ""

		err := user.Validate("createuser") //Validating the inputs/ Validando os inputs
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
				"error": err.Error(),
			})
			return
		}
		user.Password, err = models.BeforeSave(user.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
				"error": err.Error(),
			})
			return
		}
		//Saving the new User on the database/ Salvando o novo usuario na base de dados
		if dbc := db.Create(&user); dbc.Error != nil { //Return the error by JSON / Retornando o erro por JSON
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": dbc.Error})
			return
		} //Return the post data if is ok, by JSON/ Retornando o que foi postado se tudo ocorreu certo
		if err := SendConfirmationCreateAccountEmail(user.Email, c); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"waiting": "Verifique seu email !",
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "nickname em uso"})
	return

}
func ConfirmUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var UserToken models.UserAccountBadListToken
	if err := db.Where("token = ?", c.Query("t")).First(&UserToken).Error; err != nil {
		UserToken.Token = c.Query("t")
		if err := db.Create(&UserToken).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if err := authc.ValidateToken(c.Query("t")); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		var user models.User
		if err := db.Where("email = ?", c.Query("e")).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Registro não encontrado",
			})
			return
		}
		if err := db.Model(&user).Update("actived", true).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		//Here i'm returnin the message of account confirmated !/ Aqui estou retornando a mensagem que a conta foi confirmada ! :D
		//Future plains style with HTML and CSS, maybe some javascript too/ Planos futuros, estilizar com HTML e CSS, talvez um javascript junto
		c.JSON(http.StatusOK, gin.H{"success": "Conta confirmada ! "})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "link já utilizado"})
	return
}
func ChangeExternalPassword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var token models.UserAccountBadListToken
	token.Token = c.Query("t")
	if err := authc.ValidateToken(token.Token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Alguma coisa não deu certo, por favor, requiste novamente a recuperação de senha"})
		return
	}
	var input models.UserInputRecoveryPassword
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Alguma coisa deu errado",
		})
		return
	}
	if input.Password == "" || input.ConfirmPassword == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Verifique senha !"})
		return
	}
	if input.Password != input.ConfirmPassword {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "As senhas não conferem !"})
		return
	}
	//(hashadpassword,password),
	//hashad = crypted password, password is the normal one/ hashadpassword = é a senha cryptografada, passoword é a senha normal
	if err := security.VerifyPassword(user.Password, input.Password); err != nil {
		if err := models.PasswordCheck(input.Password); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}
		input.Password, err = models.BeforeSave(input.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
				"error": err.Error(),
			})
			return
		}
		if dbc := db.Create(&token); dbc.Error != nil { //Return the error by JSON / Retornando o erro por JSON
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": dbc.Error.Error()})
			return
		}
		db.Model(&user).Update("password", input.Password)
		c.JSON(http.StatusOK, gin.H{"succes": "senha alterada ! "})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "A senha não pode ser a mesma que a anterior !"})
	return
}
func RecoveryPasswordUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	if c.Query("e") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Preencha o campo email"})
		return
	}
	var user models.User
	if err := db.Where("LOWER(email) = LOWER(?)", c.Query("e")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ //We do this, to avoid user enumaration
			//Fazemos isso para evitar enumeração de usuario
			"success": "email enviado !",
		})
		return
	}
	userEmail := c.Query("e")
	userRecoveryPassword, err := authc.GenerateJWT(userEmail, 600)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	userURL := "http://localhost:8080/user/password/recovery?e=" + userEmail + "&t=" + userRecoveryPassword
	//Thanks https://blog.mailtrap.io/golang-send-email/#Sending_emails_with_smtpSendMail for the code
	msg := []byte("To: " + userEmail + "\r\n" +
		"Subject: NÂO RESPONDA ESTE EMAIL - RECUPERAÇÃO DE SENHA \r\n" +
		"\r\n" +
		"Utilize o link abaixo para recuperar a sua senha, ele é valido nos proximos 10 minutos \r\n" +
		"\r\n" +
		"Clique aqui para confirmar sua conta = " + userURL +
		"\r\n" +
		"Caso não consiga, é só copiar o link e colar no navegador ! " +
		"\r\n" +
		"Caso não tenha requistado, sugerimos trocar a senha imediatamente !")

	if err := SendEmail(userEmail, msg); err != nil {
		//Declaring and inicializing a userAccountCreatedToken
		//Declarando e inicializando uma variavel userAccount
		userAccountBadListToken := models.UserAccountBadListToken{
			Token: userRecoveryPassword,
		}
		db := c.MustGet("db").(*gorm.DB) //Establish conection with database/Estabelecendo conexão com o banco de dados
		if err := db.Create(userAccountBadListToken).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": "cheque seu email !"})
	return
}
func CreateProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("nickname = ?", c.Query("nickname")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "nickname not found",
		})
		return
	}
	if user.Profile.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "user already have a profile",
		})
		return
	}
	var input models.ProfileInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	profile := models.Profile{
		IDUser:         user.ID,
		Avatar:         input.Avatar,
		DataNascimento: input.DataNascimento,
	}
	if err := db.Create(&profile).Error; err != nil { //Return the error by JSON / Retornando o erro por JSON
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": profile})
	return
}
func UpdateUser(c *gin.Context) {
	handlers.UpdateUser(c)
}
func FindAllUsers(c *gin.Context) {
	handlers.FindUser(c)
}
func FindUserByNick(c *gin.Context) {
	handlers.FindUserByNick(c)
}
func SoftDeletedUserByNick(c *gin.Context) {
	handlers.SoftDeletedUserByNick(c)
}
