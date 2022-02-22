package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/library/cache"
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
