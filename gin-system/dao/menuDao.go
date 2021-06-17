package dao

import "gin-system/models"

func GetMenuByLevel(level int)  (error, []models.Menu) {
	var menuData []models.Menu
	err := db.Table("sys_menu").Where("menu_level = ?", level).Find(&menuData).Error
	return err, menuData
}
