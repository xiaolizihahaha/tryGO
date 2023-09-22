package router

import (
	"test1/controller/api1"
	"test1/controller/api2"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	router := gin.Default()

	router.GET("api1/add", api1.Add)
	router.GET("api1/hello", api1.Hello)
	router.GET("api2/cookie", api2.Cookie)
	router.GET("api2/header", api2.Header)
	router.GET("api2/login", api2.Login)
	router.GET("api2/home", api2.Authorize, api2.Home)

	router.POST("api1/userquery", api1.GetUser)
	router.POST("api1/queryone", api1.Select1)
	router.GET("api1/queryall", api1.Selectall)

	router.POST("api1/useradd", api1.AddUser)
	router.POST("api1/usersadd", api1.AddUsers)

	router.POST("api1/userdelete", api1.DeleteUser)
	router.POST("api1/userupdate", api1.Userupdate)

	router.POST("api1/uploadfile", api1.SaveFile)

	return router
}
