package dao

import (
	"gin-system/models"
	"gin-system/pkg/gredis"
	"strconv"
)

func GetMenuByLevel(level int) (error, []models.Menu) {
	var menuData []models.Menu
	cache := "mksys:menu" + strconv.Itoa(level)
	f, err :=gredis.RedisConn.Exists(cache)
	if f {
		gredis.RedisConn.GetObject(cache, &menuData)
		return nil, menuData
	}
	err = db.Table("sys_menu").Where("menu_level = ?", level).Find(&menuData).Error
	if len(menuData) != 0 {
		gredis.RedisConn.Set(cache, menuData, 10000)
	}
	return err, menuData
}
