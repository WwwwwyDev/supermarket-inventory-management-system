package services

import (
	"gin-system/dao"
	"gin-system/models"
)

func GetAllStaff(staffParam map[string]interface{}) (error, []models.Staff, int64) {
	err, staffData, total := dao.GetAllStaff(staffParam)
	return err, staffData, total
}

func GetStaffByName(name string) (error, []models.Staff, int64) {
	var staffData []models.Staff
	page := 1
	pageSize := 1
	searchName := name
	err, staffData, total := dao.GetAllStaff(map[string]interface{}{"page": page, "limit": pageSize, "searchName": searchName})
	return err, staffData, total
}

func AddStaff(staff models.Staff) (error) {
	err := dao.AddStaff(staff)
	return err
}

func UpdateStaff(staff models.Staff) (error) {
	err := dao.UpdateStaff(staff)
	return err
}

func DelStaff(id int) error {
	err := dao.DelStaff(id)
	return err
}

func GetStaffById(id int) (error, models.Staff) {
	err, staffData:= dao.GetStaffById(id)
	return err, staffData
}
