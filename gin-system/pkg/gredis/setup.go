package gredis

import "gin-system/pkg/setting"

var RedisConn *Cacher

func Setup()  {
	options := Options{
		MaxIdle:setting.RedisSetting.MaxIdle,
		IdleTimeout:setting.RedisSetting.IdleTimeout,
	}
	RedisConn,_= New(options)
}

