package models

import "time"

type Staff struct {
	StaffId        int       `gorm:"column:staff_id"`
	StaffName      string    `gorm:"column:staff_name"`
	StaffPassword  string    `gorm:"column:staff_password"`
	StaffLevel     int    `gorm:"column:staff_level"`
	StaffTelephone string    `gorm:"column:staff_telephone"`
	StaffSalary    int64    `gorm:"column:staff_salary"`
	StaffRemarks   string    `gorm:"column:staff_remarks"`
	IsDel          bool      `gorm:"column:is_del"`
	UpdateTime     time.Time `gorm:"column:update_time"`
}
