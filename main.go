package main

import (
	"aaa311/utils"
	"aaa311/views"

	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initialize()
	router := getEngine()
	router.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
}

//getEngine sets all the backend endpoints and handlers
func getEngine() *gin.Engine {
	router := gin.Default()
	router.Use(gin.ErrorLoggerT(gin.ErrorTypePrivate))

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.POST("/authenticate", views.Authenticate)

	router.POST("/sign-up", views.CreateUser)

	router.POST("/recovery/password", views.ForgotPassword)
	router.POST("/recovery/change-password", views.ChangePassword)

	v1 := router.Group("/api/aaa/v1", utils.CheckJWTToken())
	{
		v1.POST("/verify", views.Verify)
	}

	users := v1.Group("/users")
	{
		users.GET("/:username/roles", views.GetUserRoles)
		users.POST("/:username/roles", views.AssociateRoles)
		users.DELETE("/:username/roles/:role", views.DeleteRoles)
		users.PUT("/:username", views.UpdateUser)
	}
	return router
}

func initialize() {
	//read the config file and sets various variables used throughout the backend
	viper.AddConfigPath("/run/secrets")
	viper.AddConfigPath("/opt")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.BindEnv("database.dir", "DATABASE_DIR")
	viper.BindEnv("mail.sender", "MAIL_SENDER")
	viper.BindEnv("mail.smtp", "MAIL_SMTP")
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetLevel(log.DebugLevel)

	//sets constants used by the token related functions
	utils.InitiateTokenParams()
}
