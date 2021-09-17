package config

import (
	"flag"
	"fmt"
	"time"
)

var (
	//HTTP 监听的端口号
	HTTPPort uint = 3001
	//Mode 运行模式
	Mode = "dev"
	//Mysql 数据库信息
	Mysql mysqlConf
	//Redis 数据库信息
	Redis redisConf
	//TokenName 前台设置cookie名称
	TokenName = "JAVASESSID"
	//FrontEnd 前端域名列表
	FrontEnd []string
	//AdminDomain 后端域名
	AdminDomain = "http://192.168.2.100:3001"
	//TimeFormat 格式化时间
	TimeFormat = "2006-01-02 15:04:05"
	//Limit 分页一页显示数量
	Limit = 20
	//PathRoot 项目静态资源目录
	PathRoot = "/"

	Code = codeConf{1, 0}
)

//开发模式
func dev() {
	Mysql = mysqlConf{
		Host:         "127.0.0.1",
		Database:     "test",
		UserName:     "root",
		Password:     "root",
		Port:         3306,
		MaxIdleConns: 1000,
		MaxOpenConns: 2000,
	}
	Redis = redisConf{
		Host:        "192.168.2.222",
		Password:    "",
		Port:        "6601",
		MaxIdle:     1000,
		MaxActive:   0,
		IdleTimeout: 300,
		MaxTimeout:  time.Duration(30) * time.Second,
	}
}

//生产模式
func produce() {
	Mode = "produce"
	Mysql = mysqlConf{
		Host:         "127.0.0.1",
		Database:     "go_produce",
		UserName:     "root",
		Password:     "root",
		Port:         3306,
		MaxIdleConns: 1000,
		MaxOpenConns: 3000,
	}
	Redis = redisConf{
		Host:        "127.0.0.1",
		Password:    "",
		Port:        "6379",
		MaxIdle:     1000,
		MaxActive:   0,
		IdleTimeout: 300,
		MaxTimeout:  time.Duration(30) * time.Second,
	}
	FrontEnd = []string{
		"http://localhost",
	}
}

//初始化配置
func Run() {
	var mode string

	//设置模式
	flag.StringVar(&mode, "mode", "dev", "运行模式")
	flag.Parse()

	if mode == "produce" {
		produce()
	} else {
		dev()
	}
	fmt.Println("设置初始化成功，应用：" + mode + "模式")
}

type mysqlConf struct {
	Host         string
	Database     string
	UserName     string
	Password     string
	Port         int
	MaxIdleConns int /*用于设置闲置的连接数。*/
	MaxOpenConns int /*用于设置最大打开的连接数，默认值为0表示不限制*/
}

type redisConf struct {
	Host        string
	Password    string
	Port        string
	MaxIdle     int           /*用于设置闲置的连接数。*/
	MaxActive   int           /*连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配*/
	IdleTimeout time.Duration /*连接关闭时间 300秒 （300秒不使用自动关闭）*/
	MaxTimeout  time.Duration
}

type codeConf struct {
	Ok int
	No int
}
