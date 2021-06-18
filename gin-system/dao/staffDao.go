package dao

import (
	"gin-system/models"
)

func GetAllStaff(staffParam map[string]interface{}) (error, []models.Staff, int64) {
	var staffData []models.Staff
	page := staffParam["page"].(int)
	pageSize := staffParam["limit"].(int)
	searchName := staffParam["searchName"].(string)
	var total int64
	err := db.Table("erp_staff").Where("staff_name like ? and is_del = false", "%"+searchName+"%").Order("staff_id ASC").Count(&total).Offset((page-1)*pageSize).Limit(pageSize).Find(&staffData).Error
	return err, staffData, total
}

func GetStaffById(id int) (error, models.Staff) {
	var staffData models.Staff
	err := db.Table("erp_staff").Where("staff_id = ? and is_del = false", id).Find(&staffData).Error
	return err, staffData
}

func AddStaff(staff models.Staff) (error) {
	err := db.Table("erp_staff").Select("staff_name", "staff_password", "staff_level", "staff_telephone", "staff_salary", "staff_remarks", "is_del",
		"update_time").Create(&staff).Error
	return err
}
//
func UpdateStaff(staff models.Staff) (error) {
	err := db.Table("erp_staff").Where("staff_id = ?", staff.StaffId).Updates(&staff).Error
	return err
}

func DelStaff(id int) error {
	err := db.Table("erp_staff").Where("staff_id = ?", id).Update("is_del", true).Error
	return err
}
