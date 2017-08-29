package views

import (
	"aaa311/database"
	"aaa311/models"
	"errors"
	"net/http"

	gomail "gopkg.in/gomail.v2"

	"aaa311/utils"

	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type recoverySchema struct {
	Username string
	Password string
}

//ForgotPassword creates a new random password and sends the password to the user email
func ForgotPassword(c *gin.Context) {
	var user recoverySchema
	var fullUserInfo models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad JSON"))
		return
	}

	log.Infof("bind user: %+v", user)

	db := database.InitiateConnection()

	err = db.Read("user", user.Username, &fullUserInfo)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusNotFound, errors.New("User not found"))
		return
	}

	newPassword := utils.RandomPassword(7)

	log.Infof(fmt.Sprintf("new password %s", newPassword))

	crpytPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 10)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fullUserInfo.Password = string(crpytPassword)

	log.Infof("full user info %+v", fullUserInfo)

	err = db.Write("user", fullUserInfo.Username, fullUserInfo)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", viper.GetString("mail.sender"))
	m.SetHeader("To", fullUserInfo.Email)
	m.SetHeader("Subject", "You requested a new password!")
	m.SetBody("text/html", "your new password is: "+newPassword)

	d := gomail.Dialer{Host: viper.GetString("mail.smtp"), Port: 25}
	if err := d.DialAndSend(m); err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Failed to send email"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "password changed"})
}

//ChangePassword changes the user password with a new password provided by the user
func ChangePassword(c *gin.Context) {
	var user recoverySchema
	var fullUserInfo models.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad JSON"))
		return
	}
	log.Infof("bind user: %+v", user)

	db := database.InitiateConnection()

	err = db.Read("user", user.Username, &fullUserInfo)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Infof("user info: %+v", fullUserInfo)

	cryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fullUserInfo.Password = string(cryptPassword)

	err = db.Write("user", fullUserInfo.Username, fullUserInfo)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "password changed"})
}
