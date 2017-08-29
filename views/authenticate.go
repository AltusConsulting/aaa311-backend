package views

import (
	"time"

	"aaa311/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// AuthenticationSchema Test
type AuthenticationSchema struct {
	Username string
	Password string
}

// Authenticate checks if the username and password are correct and generates a token
func Authenticate(c *gin.Context) {
	var requestUser AuthenticationSchema
	var token map[string]string
	token = make(map[string]string)

	err := c.BindJSON(&requestUser)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad request"})
		return
	}

	user, err := utils.AuthenticateUser(requestUser.Username, requestUser.Password)
	if err != nil {
		log.Error(err)
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	claims := make(map[string]interface{})
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["username"] = user.Username
	claims["first_name"] = user.FirstName
	claims["last_name"] = user.LastName
	claims["email"] = user.Email
	claims["roles"] = user.Roles
	tokenString, err := utils.GenerateToken(claims)

	if err != nil {
		c.JSON(500, gin.H{"message": "Could not generate token"})
		log.Warning(err)
		return
	}

	token["token"] = tokenString
	c.JSON(200, token)
}
