package routers

import (
	_ "demoProject/docs"
	"demoProject/middleware"
	"demoProject/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.Ping)
	r.GET("/problemList", service.GetProblemList)
	r.GET("/problemDetail", service.GetProblemDetail)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 用户详情查找
	r.GET("/userDetail", middleware.AuthAdminCheck(), service.GetUserDetail)

	// 用户登录
	r.POST("/userLogin", service.UserLogin)

	return r

}
