package views

import (
	"aaa311/database"
	"aaa311/models"
	"errors"

	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//CreateUser creates a new user in the scribble database
func CreateUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad JSON"))
		return
	}

	if userExist(user.Username) {
		log.Error("user already exists")
		c.AbortWithError(http.StatusBadRequest, errors.New("user already exists"))
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.Password = string(password)

	db := database.InitiateConnection()
	err = db.Write("user", user.Username, user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New("Failed to create user"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "user created"})
}

//UpdateUser updates a user info in the scribble database
func UpdateUser(c *gin.Context) {
	userName := c.Param("username")
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad JSON"))
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.Password = string(password)

	db := database.InitiateConnection()
	err = db.Write("user", userName, user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusInternalServerError, errors.New("Failed to update User"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "user updated"})
}

// GetUserRoles Get roles by user
func GetUserRoles(c *gin.Context) {
	userName := c.Param("username")
	var user models.User

	db := database.InitiateConnection()

	err := db.Read("user", userName, &user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user.Roles)
}

// AssociateRoles ...
func AssociateRoles(c *gin.Context) {
	userName := c.Param("username")
	var user models.User
	var roles []models.Roles

	err := c.BindJSON(&roles)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("Bad JSON"))
		return
	}

	db := database.InitiateConnection()

	err = db.Read("user", userName, &user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user.Roles = roles

	err = db.Write("user", userName, user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "roles assigned to user " + userName})
}

// DeleteRoles Delete role by user
func DeleteRoles(c *gin.Context) {
	userName := c.Param("username")
	role := c.Param("role")
	var user models.User

	db := database.InitiateConnection()

	err := db.Read("user", userName, &user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var index int
	for k, v := range user.Roles {
		if v.Name == role {
			index = k
			return
		}
	}

	newRoles := remove(user.Roles, index)
	user.Roles = newRoles

	err = db.Read("user", userName, user)
	if err != nil {
		log.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "role deleted"})
}

func remove(s []models.Roles, i int) []models.Roles {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

//userExist checks if the user already exist in the database
func userExist(username string) bool {
	var user models.User
	db := database.InitiateConnection()
	err := db.Read("user", username, &user)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}
