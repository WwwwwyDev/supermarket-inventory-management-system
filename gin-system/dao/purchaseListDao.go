package dao

import "gin-system/models"

func GetAllPurchaseList(purchaseListParam map[string]interface{}) (error, []models.PurchaseList, int64) {
	var purchaseListData []models.PurchaseList
	page := purchaseListParam["page"].(int)
	pageSize := purchaseListParam["limit"].(int)
	searchName := purchaseListParam["searchName"].(string)
	var total int64
	err := db.Table("erp_purchase_list p").Select("p.purchase_list_id, p.purchase_list_number, p.purchase_list_price, p.purchase_list_remarks, p.purchase_list_staff, s.staff_name purchase_list_staff_name, p.purchase_list_time, p.update_time, p.is_del").Joins("INNER JOIN erp_staff s ON p.purchase_list_staff = s.staff_id").Where("s.staff_name like ? and p.is_del = false", searchName+"%").Order("purchase_list_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&purchaseListData).Error
	return err, purchaseListData, total
}

func DelPurchaseList(id int) error {
	err := db.Table("erp_purchase_list").Where("purchase_list_id = ?", id).Update("is_del", true).Error
	return err
}

func UpdatePurchaseList(purchaseList models.PurchaseList) error {
	err := db.Table("erp_purchase_list").Select("purchase_list_staff",  "purchase_list_remarks", "update_time", "is_del").Where("purchase_list_id = ?", purchaseList.PurchaseListId).Updates(&purchaseList).Error
	return err
}

func AddPurchaseList(purchaseList models.PurchaseList) (error) {
	err := db.Table("erp_purchase_list").Select("purchase_list_staff", "purchase_list_number", "purchase_list_price", "purchase_list_time", "purchase_list_remarks", "update_time", "is_del").Create(&purchaseList).Error
	return err
}