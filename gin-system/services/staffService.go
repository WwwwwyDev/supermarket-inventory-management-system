package services

import (
	"gin-system/dao"
	"gin-system/models"
	"time"
)

func GetAllStaff(staffParam map[string]interface{}) (error, []models.Staff, int64) {
	var staffData []models.Staff
	page := staffParam["page"].(int)
	pageSize := staffParam["limit"].(int)
	searchName := staffParam["searchName"].(string)
	err, staffData, total := dao.GetAllStaff(map[string]interface{}{"page": page, "limit": pageSize, "searchName": searchName})
	return err, staffData, total
}

func AddStaff(staff models.Staff) (error, models.Staff) {
	err := dao.AddStaff(staff)
	return err, staff
}

func UpdateStaff(staff models.Staff) (error, models.Staff) {
	staff.UpdateTime = time.Now()
	err := dao.UpdateStaff(staff)
	return err, staff
}

func DelStaff(id int) error {
	err := dao.DelStaff(id)
	return err
}


