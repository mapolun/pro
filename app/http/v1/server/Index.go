package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/common/response"
	"pro/app/database"
)

func Index(c *gin.Context) {
	redis := database.RedisInter.Get()
	r, err := redis.Do("Set", "test", 111)
	if err != nil {
		fmt.Println(err)
	}
	response.Success(c, "ok", r)
}
