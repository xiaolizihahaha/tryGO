package api1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test1/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	//url
	// var u models.User
	// id1 := c.Query("id")
	// id2, err := strconv.Atoi(id1)
	// if err != nil {
	// 	fmt.Println("id1=", id1)
	// 	fmt.Println("id2=", id2)
	// 	fmt.Println(err.Error())
	// 	id2 = 2
	// }

	//body
	u := models.User{}
	c.BindJSON(&u)
	id2 := u.User_id

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"user": u.Select(int64(id2)),
	})
}

func AddUser(c *gin.Context) {
	//get url query
	// name := c.Query("user_name")
	// password := c.Query("user_password")
	// u := models.User{User_name: name, User_password: password}

	//post bodyJson bind
	u := models.User{}
	c.BindJSON(&u)
	// fmt.Println(u)

	id := u.AddUser()
	message := fmt.Sprintf("insert user %d success", id)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
	})
}

func DeleteUser(c *gin.Context) {

	//url
	// var u models.User
	// id1 := c.Query("id")
	// id2, err := strconv.Atoi(id1)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	id2 = 1
	// }

	//body
	u := models.User{}
	c.BindJSON(&u)

	id2 := u.User_id
	u.Delete(int64(id2))

	message := fmt.Sprintf("delete user %d success", id2)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
	})
}

func Userupdate(c *gin.Context) {

	//url
	// username := c.Query("username")
	// password := c.Query("password")
	// id1 := c.Query("id")
	// id2, err := strconv.Atoi(id1)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	id2 = 1
	// }

	// u := models.User{User_id: int64(id2), User_name: username, User_password: password}

	//body
	u := models.User{}
	c.BindJSON(&u)

	id2 := u.User_id

	u.Update()
	message := fmt.Sprintf("update user %d success", id2)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
	})
}

func Select1(c *gin.Context) {
	u := models.User{}
	c.BindJSON(&u)

	result := u.SelectOne(int64(u.User_id))

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"user": result,
	})

}

func Selectall(c *gin.Context) {
	u := models.SelectAll()

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"users": u,
	})
}

func AddUsers(c *gin.Context) {
	usersList := c.PostForm("userList")
	var users []models.User

	err := json.Unmarshal([]byte(usersList), &users)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"code":     200,
			"messsage": "userList error",
		})
		return
	}

	saveSign := models.SaveUsers(users)
	if saveSign == true {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "save success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":     200,
			"messsage": "save failed",
		})
	}

}

func SaveFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "upload file failed",
		})
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err1 := c.SaveUploadedFile(file, "file/"+file.Filename)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "upload file failed",
		})
		c.AbortWithError(http.StatusBadRequest, err1)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload file success",
	})

}
