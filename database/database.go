package database

import (
	log "github.com/Sirupsen/logrus"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/spf13/viper"
)

var DB *scribble.Driver

//InitiateConnection starts the connection to the database
//and returns a scribble driver to use the database functions
func InitiateConnection() *scribble.Driver {
	DB, err := scribble.New(viper.GetString("database.dir"), nil)
	if err != nil {
		log.Fatal(err)
	}
	return DB
}
