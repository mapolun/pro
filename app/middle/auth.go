package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro/config"
)

func Access(c *gin.Context) {
	var origin = c.Request.Header.Get("Origin")
	if config.Mode != "dev" {
		var isOk = false
		isOk = false
		allow := config.FrontEnd
		for _, v := range allow {
			if origin == v {
				isOk = true
				break
			}
		}

		if isOk != true {
			c.String(http.StatusForbidden, "服务内部错误，请联系管理人员！")
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
