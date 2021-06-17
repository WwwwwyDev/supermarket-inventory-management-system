package main

import (
	"fmt"
	"gin-system/dao"
	"gin-system/models"
	"gin-system/pkg/setting"
	"time"
)

func init() {
	setting.Setup()
	dao.Setup()
}

func main() {

	c := models.Staff{12, "wwy", "test78", 2, "34213", 123124, "test", false, time.Now()}
	dao.UpdateStaff(c)
	fmt.Println(dao.GetAllStaff(map[string]interface{}{"page": 1, "limit": 10000, "searchName": "wwy"}))
}
