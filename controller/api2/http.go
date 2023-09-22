package api2

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cookie(c *gin.Context) {
	cookie, err := c.Cookie("cookieName")
	if err != nil {
		cookie = "hello"
		c.SetCookie("cookieName", cookie, 60*2, "/", "localhost", false, true)
		message := fmt.Sprintf("cookie have changed to %s", cookie)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": message,
		})
		return
	}
	message := fmt.Sprintf("cookie have't changed (%s)", cookie)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
	})

}

func Header(c *gin.Context) {
	fmt.Println("---header/--- rn")
	// fmt.Println(c.Request.Header)
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	header := c.Request.Header
	cookie := c.Request.Header["Cookie"]

	c.JSON(http.StatusOK, gin.H{
		"header": header,
		"cookie": cookie,
	})

}
