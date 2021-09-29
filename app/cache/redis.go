package cache

import (
	"github.com/garyburd/redigo/redis"
	"pro/config"
)

var RedisInter *redis.Pool

//创建redis连接池
func RedisInit() error {
	pool := &redis.Pool{
		MaxIdle:     config.Redis.MaxIdle,
		MaxActive:   config.Redis.MaxActive,
		IdleTimeout: config.Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			con, err := redis.Dial(
				"tcp",
				config.Redis.Host+":"+config.Redis.Port, // address
				redis.DialPassword(config.Redis.Password),
				redis.DialConnectTimeout(config.Redis.MaxTimeout),
				redis.DialReadTimeout(config.Redis.MaxTimeout),
				redis.DialWriteTimeout(config.Redis.MaxTimeout),
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
