package utils

import (
	"aaa311/database"
	"aaa311/models"

	"golang.org/x/crypto/bcrypt"
)

// AuthenticateUser is ...
func AuthenticateUser(username string, password string) (models.User, error) {
	var user models.User

	db := database.InitiateConnection()
	err := db.Read("user", username, &user)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
