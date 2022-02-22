package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"pro/app/common/fn"
	"pro/app/library/response"
	"pro/config"
	"strings"
	"time"
)

const imageSize = 4096

func UploadImage(c *gin.Context) {
	f, err := c.FormFile("image")
	if err != nil {
		response.Error(c, "上传失败")
		return
	}

	fileExt := strings.ToLower(path.Ext(f.Filename))

	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		response.Error(c, "文件类型不合法")
		return
	}

	if f.Size > imageSize*1024 {
		response.Error(c, "文件过大")
		return
	}

	//上传文件名
	fileName := fn.Md5(fmt.Sprintf("%s%s", f.Filename, time.Now().String())) + fileExt

	//获取上传目录
	uploadTmpDir := fmt.Sprintf("%d-%d/%d", time.Now().Year(), time.Now().Month(), time.Now().Day())

	uploadDir := fmt.Sprintf("%s/%s", config.Get.Upload.Dir, uploadTmpDir)

	if fn.DirExists(uploadDir) == false {
		_ = os.MkdirAll(uploadDir, os.ModePerm)
	}

	filePath := fmt.Sprintf("%s/%s", uploadDir, fileName)

	err = c.SaveUploadedFile(f, filePath)
	if err != nil {
		log.Println("图片错误：", err.Error())
		response.Error(c, "图片上传错误")
		return
	}

	newFilePath := fmt.Sprintf("/upload/%s/%s", uploadTmpDir, fileName)

	response.Success(c, "上传成功", newFilePath)
}
