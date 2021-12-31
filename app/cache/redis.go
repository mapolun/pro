package cache

import (
	"github.com/garyburd/redigo/redis"
	"pro/config"
)

var RedisInter *redis.Pool

//创建redis连接池
func RedisInit() error {
	pool := &redis.Pool{
		MaxIdle:     config.Get.Redis.MaxIdle,
		MaxActive:   config.Get.Redis.MaxActive,
		IdleTimeout: config.Get.Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			con, err := redis.Dial(
				"tcp",
				config.Get.Redis.Host+":"+config.Get.Redis.Port, // address
				redis.DialPassword(config.Get.Redis.Password),
				redis.DialConnectTimeout(config.Get.Redis.MaxTimeout),
				redis.DialReadTimeout(config.Get.Redis.MaxTimeout),
				redis.DialWriteTimeout(config.Get.Redis.MaxTimeout),
			)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	RedisInter = pool
	return nil
}
