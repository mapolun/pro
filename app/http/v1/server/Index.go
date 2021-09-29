package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/cache"
	"pro/app/common/response"
)

func Index(c *gin.Context) {
	redis := cache.RedisInter.Get()
	r, err := redis.Do("Set", "test", 111)
	if err != nil {
		fmt.Println(err)
	}
	response.Success(c, "ok", r)
}

func Test(c *gin.Context) {
	type params struct {
		Id   int    `json:"id"   binding:"required"`
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBind(&params{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
	}
	fmt.Println(params{})

	//dec := json.NewDecoder(strings.NewReader(jsonstring))
	//fmt.Println(dec)
}
