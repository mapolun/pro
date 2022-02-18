package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro/config"
)

//Success 正确输出
func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    config.Get.Code.Ok,
		"message": msg,
		"data":    data,
	})
	c.Abort()
}

//Error 错误输出
func Error(c *gin.Context, msg string, args ...int) {
	c.JSON(http.StatusOK, gin.H{
		"code":    config.Get.Code.No,
		"message": msg,
	})
	c.Abort()
}

//Handler 自定义输出
func Handler(c *gin.Context, msg string, code int, data interface{}) {
	h := gin.H{
		"code":    code,
		"message": msg,
	}

	if data != nil {
		h["data"] = data
	}
	c.JSON(code, h)
	c.Abort()
}
