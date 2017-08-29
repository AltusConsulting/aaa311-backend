package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//test handler that runs once the token has been verified
func Verify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mssg": "verify"})
}
