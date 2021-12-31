package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

var Get *config

type config struct {
	HttpPort   uint     //http服务监听端口
	Mode       string   //模式 dev=开发模式 produce=生产模式
	FrontEnd   []string //前端域名列表
	Limit      int      //分页一页显示数量
	TimeFormat string   //格式化时间
	Upload     *upload
	Mysql      *mysqlConf
	Redis      *redisConf
	Code       *codeConf
}

//文件配置
type upload struct {
	Host string
	Dir  string
}

//mysql
type mysqlConf struct {
	Host         string
	Database     string
	UserName     string
	Password     string
	Port         int
	MaxIdleConns int /*用于设置闲置的连接数。*/
	MaxOpenConns int /*用于设置最大打开的连接数，默认值为0表示不限制*/
}

//redis
type redisConf struct {
	Host        string
	Password    string
	Port        string
	MaxIdle     int           /*用于设置闲置的连接数。*/
	MaxActive   int           /*连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配*/
	IdleTimeout time.Duration /*连接关闭时间 300秒 （300秒不使用自动关闭）*/
	MaxTimeout  time.Duration
}

//公用错误码配置
type codeConf struct {
	Ok int
	No int
}

//初始化配置
func Run() {
	var mode string
	//设置模式
	flag.StringVar(&mode, "mode", "dev", "运行模式")
	conf, err := readToml()
	if err != nil {
		panic("配置文件解析失败" + err.Error())
	}
	conf.Mode = mode
	Get = conf
	fmt.Println("设置初始化成功，应用：" + mode + "模式")
}

//读取toml配置文件
func readToml() (*config, error) {
	conf := &config{}
	_, err := toml.DecodeFile("./config/config.toml", conf)
	if err != nil {
		return conf, err
	}
	return buildOtherConf(conf), nil
}

//其他配置项
func buildOtherConf(conf *config) *config {
	conf.TimeFormat = "2006-01-02 15:04:05"
	conf.Limit = 100
	conf.Code = &codeConf{0, 1}
	return conf
}
