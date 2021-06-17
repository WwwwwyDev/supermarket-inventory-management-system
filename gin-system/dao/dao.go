package dao

import (
	"fmt"
	"gin-system/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB


func Setup() {
	var err error
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}