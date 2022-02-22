package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"pro/app/common/fn"
	"pro/app/library/logger"
	"pro/app/library/response"
	"pro/app/repository/request"
	"strconv"
	"time"
)

func WechatComplain(c *gin.Context) {

	var params request.WechatComplainRequest

	if err := c.ShouldBind(&params); err != nil {
		return
	}

	check := fn.ValidateMobile(params.Mobile)
	if check == false {
		response.Error(c, "请填写正确的手机号码")
		return
	}

	for _, image := range params.Image {
		check = fn.ValidateImage(image)
		if check == false {
			response.Error(c, "请上传正确的图片格式")
			return
		}
	}

	//请求IP
	params.ClientIp = c.ClientIP()

	//转json
	data, err := json.Marshal(params)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	//存储到log内容里
	logs := new(logger.Logger)
	year := strconv.Itoa(time.Now().Year())
	month := strconv.Itoa(int(time.Now().Month()))
	day := strconv.Itoa(time.Now().Day())

	filePaths := fmt.Sprintf("/%s/%s/%s", "wechat", year+"-"+month, day+".log")

	logs.SetFilePath(filePaths)
	l, err := logs.New()
	if err != nil {
		response.Error(c, "系统错误")
		return
	}
	l.Log(logrus.InfoLevel, string(data))

	response.Success(c, "您的投诉已经提交，请等待处理", struct{}{})
}
