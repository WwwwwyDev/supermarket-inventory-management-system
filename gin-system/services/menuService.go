package services

import (
	"gin-system/dao"
	"gin-system/models"
)

func GetMenuByLevel(level int)  (error, []models.Menu) {
	err, menuData := dao.GetMenuByLevel(level)
	return err, menuData
}
