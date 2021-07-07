package gredis

import "gin-system/pkg/setting"

var RedisConn *Cacher

func Setup()  {
	options := Options{
		Addr:setting.RedisSetting.Host,
		Password:setting.RedisSetting.Password,
		MaxActive:setting.RedisSetting.MaxActive,
		MaxIdle:setting.RedisSetting.MaxIdle,
		IdleTimeout:setting.RedisSetting.IdleTimeout,
	}
	RedisConn,_= New(options)
}

