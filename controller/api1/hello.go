package api1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	param := c.PostForm("second")
	fmt.Println("it is =", param)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "hello!world",
	})
}
