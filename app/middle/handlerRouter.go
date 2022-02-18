package middle

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pro/app/library/logger"
	"pro/app/library/response"
	"pro/config"
	"runtime/debug"
	"time"
)

//鉴权校验
func HandlerAccess(c *gin.Context) {
	var origin = c.Request.Header.Get("Origin")
	if config.Get.Mode != "dev" {
		var isOk = false
		isOk = false
		allow := config.Get.FrontEnd
		for _, v := range allow {
			if origin == v {
				isOk = true
				break
			}
		}

		if isOk != true {
			response.Handler(c, "服务内部错误，请联系管理人员！", http.StatusInternalServerError, nil)
			c.Abort()
		}
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,OPTIONS") //允许的请求类型
		c.Header("Access-Control-Max-Age", "172800")
		c.JSON(http.StatusOK, "Options Request!")
		return
	}
	c.Next()
}

//系统级错误异常
func HandlerRecover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic 错误 %v\n", err)
			debug.Stack()
			response.Handler(c, "服务器内部错误", http.StatusInternalServerError, nil)
		}
	}()
}

//404
func HandlerNotFound(c *gin.Context) {
	response.Handler(c, "未找到资源", http.StatusNotFound, nil)
}

//日志服务
func HandlerLogger() gin.HandlerFunc {
	logger := new(logger.Logger)
	l, err := logger.New()
	if err != nil {
		log.Println("日志错误获取：", err)
	}

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		l.Infof("| %3d | %13v | %15s | %s | %s",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
