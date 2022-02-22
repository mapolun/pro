package router

import (
	"github.com/gin-gonic/gin"
	"pro/app/http/v1/server"
	"pro/app/middle"
	"pro/config"
)

func router(route *gin.Engine) *gin.Engine {

	v1 := route.Group("/v1")
	//遊客操作，无需登录
	visitorAPI := v1.Group("/api")
	{
		visitorAPI.POST("wechatComplain", server.WechatComplain) //假投诉
		visitorAPI.POST("uploadImage", server.UploadImage)       //图片上传
	}

	return route
}

func RouteInit() *gin.Engine {
	if config.Get.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.New()

	/************************************/
	/********** 服务中间件 ********/
	/************************************/

	//日志
	gin.DisableConsoleColor() //禁用控制台日志颜色
	if config.Get.Mode != "dev" {
		route.Use(middle.HandlerLogger())
	} else {
		route.Use(gin.Logger())
	}

	//404
	route.NoRoute(middle.HandlerNotFound)
	route.NoMethod(middle.HandlerNotFound)

	//异常
	route.Use(middle.HandlerRecover)

	//鉴权
	route.Use(middle.HandlerAccess)
	return router(route)
}
