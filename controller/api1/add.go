package api1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	a := c.Query("a")
	b := c.Query("b")

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	a_value, a_err := strconv.Atoi(a)
	b_value, b_err := strconv.Atoi(b)

	if a_err != nil || b_err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    4001,
			"message": "wrong with a or b",
		})
		return
	}

	result_value := a_value + b_value
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result_value,
	})
}
