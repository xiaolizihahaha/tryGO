package router

import (
	"test1/controller/api1"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	router := gin.Default()

	router.GET("api1/add", api1.Add)
	router.GET("api1/hello", api1.Hello)
	router.GET("api1/userquery", api1.GetUser)
	router.GET("api1/useradd", api1.AddUser)
	router.GET("api1/userdelete", api1.DeleteUser)
	router.GET("api1/userupdate", api1.Userupdate)
	return router
}
