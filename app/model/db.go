package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"pro/config"
)

var Db *gorm.DB

//gorm链接数据库
func Run() error {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
			config.Get.Mysql.UserName,
			config.Get.Mysql.Password,
			config.Get.Mysql.Host,
			config.Get.Mysql.Port,
			config.Get.Mysql.Database,
		),
	)
	if err != nil {
		return err
	}
	/*连接池信息*/
	db.DB().SetMaxIdleConns(config.Get.Mysql.MaxIdleConns) //设置最大空闲数
	db.DB().SetMaxOpenConns(config.Get.Mysql.MaxOpenConns) //设置最大连接数
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "ww_" + defaultTableName
	}
	Db = db
	return nil
}
