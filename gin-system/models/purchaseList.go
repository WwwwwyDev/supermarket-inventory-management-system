package models

import "time"

type PurchaseList struct {
	PurchaseListId        int       `gorm:"column:purchase_list_id"`
	PurchaseListStaff     int       `gorm:"column:purchase_list_staff"`
	PurchaseListStaffName string       `gorm:"column:purchase_list_staff_name"`
	PurchaseListNumber    int       `gorm:"column:purchase_list_number"`
	PurchaseListPrice     int       `gorm:"column:purchase_list_price"`
	PurchaseListTime      time.Time `gorm:"column:purchase_list_time"`
	PurchaseListRemarks   string    `gorm:"column:purchase_list_remarks"`
	UpdateTime            time.Time `gorm:"column:update_time"`
	IsDel                 bool      `gorm:"column:is_del"`
}
