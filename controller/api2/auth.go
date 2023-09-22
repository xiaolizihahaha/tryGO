package api2

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.SetCookie("cookieName", "123", 2*60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "log success,you have token",
	})
}

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "auth success,you enter home",
	})
}

func Authorize(c *gin.Context) {
	cookie, err := c.Cookie("cookieName")
	if err == nil && cookie == "123" {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "stop!you don't have been authorized",
	})
	c.Abort()
	return
}
