package logger

import (
	"github.com/sirupsen/logrus"
	"pro/app/common/fn"
	"pro/config"
	"strconv"
	"time"
)

type Logger struct {
	filePath string
}

func (o *Logger) SetFilePath(filePath string) {
	o.filePath = filePath
}

func (o *Logger) GetFilePath() string {
	return o.filePath
}

func (o *Logger) New() (*logrus.Logger, error) {
	if o.filePath == "" {
		year := strconv.Itoa(time.Now().Year())
		month := strconv.Itoa(int(time.Now().Month()))
		day := strconv.Itoa(time.Now().Day())

		tmpPath := "/" + year + "-" + month + "/"

		fileName := day + ".log"

		o.filePath = tmpPath + fileName
	}
	filePath := fn.GetRootPath() + config.Get.Log.Dir + o.filePath

	f, err := fn.GetFile(filePath)
	if err != nil {
		return nil, err
	}
	//实例化
	logger := logrus.New()
	//设置输出
	logger.SetOutput(f)
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: config.Get.TimeFormat,
	})
	return logger, nil
}
