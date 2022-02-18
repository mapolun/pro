package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"pro/app/library/cache"
	"pro/app/library/logger"
	"pro/app/library/response"
)

func Index(c *gin.Context) {
	redis := cache.RedisInter.Get()
	r, err := redis.Do("Set", "test", 111)
	if err != nil {
		fmt.Println(err)
		return
	}
	response.Success(c, "ok", r)
}

func Test(c *gin.Context) {

	logger := new(logger.Logger)
	l, err := logger.New()
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	l.Log(logrus.InfoLevel, "123")
	l.Debug("调试")

	type paramsMod struct {
		Id   int    `form:"id" binding:"required"`
		Name string `form:"name" binding:"required"`
	}
	var params paramsMod
	if err := c.BindQuery(&params); err != nil {
		log.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	response.Success(c, "asd", nil)
}
