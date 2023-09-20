package api1

import (
	"fmt"
	"net/http"
	"strconv"
	"test1/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var u models.User
	id1 := c.Query("id")

	id2, err := strconv.Atoi(id1)
	if err != nil {
		fmt.Println("id1=", id1)
		fmt.Println("id2=", id2)
		fmt.Println(err.Error())
		id2 = 2
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"user": u.Select(int64(id2)),
	})
}

func AddUser(c *gin.Context) {
	name := c.Query("username")
	password := c.Query("password")

	u := models.User{User_name: name, User_password: password}

	id := u.AddUser()
	message := fmt.Sprintf("insert user %d success", id)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
	})
}

func DeleteUser(c *gin.Context) {
	var u models.User

	id1 := c.Query("id")
	id2, err := strconv.Atoi(id1)
	if err != nil {
		fmt.Println(err.Error())
		id2 = 1
	}
	u.Delete(int64(id2))

	message := fmt.Sprintf("delete user %d success", id2)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
	})
}

func Userupdate(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	id1 := c.Query("id")

	id2, err := strconv.Atoi(id1)

	if err != nil {
		fmt.Println(err.Error())
		id2 = 1
	}

	u := models.User{User_id: int64(id2), User_name: username, User_password: password}

	u.Update()
	message := fmt.Sprintf("delete user %d success", id2)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
	})
}
