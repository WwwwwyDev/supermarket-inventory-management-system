package models

type Menu struct {
	MenuId        int       `gorm:"column:menu_id"`
	MenuName      string    `gorm:"column:menu_name"`
	MenuAddress   string    `gorm:"column:menu_address"`
	MenuLevel     int    `gorm:"column:menu_level"`
}
