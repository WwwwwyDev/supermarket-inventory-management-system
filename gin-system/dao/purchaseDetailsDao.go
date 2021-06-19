package dao

import (
	"gin-system/models"
)

func GetAllPurchaseDetails(purchaseDetailsParam map[string]interface{}) (error, []models.PuchaseDetails, int64) {
	var purchaseDetailsData []models.PuchaseDetails
	page := purchaseDetailsParam["page"].(int)
	pageSize := purchaseDetailsParam["limit"].(int)
	searchList := purchaseDetailsParam["searchList"].(int)
	searchGoodsName := purchaseDetailsParam["searchGoodsName"].(string)
	var total int64
	if searchList == -1{
		err := db.Table("erp_purchase_details d").Select("d.purchase_details_id, d.purchase_details_list, d.purchase_details_goods, g.goods_name purchase_details_goods_name,d.purchase_details_number, d.purchase_details_price, d.purchase_details_remarks, d.update_time, d.is_del").Joins("INNER JOIN erp_goods g ON d.purchase_details_goods = g.goods_id").Where("d.is_del = false").Where("g.goods_name like ?",searchGoodsName+"%").Order("purchase_details_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&purchaseDetailsData).Error
		return err, purchaseDetailsData, total
	}
	err := db.Table("erp_purchase_details d").Select("d.purchase_details_id, d.purchase_details_list, d.purchase_details_goods, g.goods_name purchase_details_goods_name,d.purchase_details_number, d.purchase_details_price, d.purchase_details_remarks, d.update_time, d.is_del").Joins("INNER JOIN erp_goods g ON d.purchase_details_goods = g.goods_id").Where("d.purchase_details_list = ?", searchList).Where("d.is_del = false").Where("g.goods_name like ?",searchGoodsName+"%").Order("purchase_details_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&purchaseDetailsData).Error
	return err, purchaseDetailsData, total
}

func DelPurchaseDetails(id int) error {
	err := db.Table("erp_purchase_details").Where("purchase_details_id = ?", id).Update("is_del", true).Error
	return err
}

func UpdatePurchaseDetails(puchaseDetails models.PuchaseDetails) error {
	err := db.Table("erp_purchase_details").Where("purchase_details_id = ?", puchaseDetails.PurchaseDetailsId).Updates(&puchaseDetails).Error
	return err
}

func AddPurchaseDetails(puchaseDetails models.PuchaseDetails) (error) {
	err := db.Table("erp_purchase_details").Select("purchase_details_list", "purchase_details_goods", "purchase_details_number", "purchase_details_price", "purchase_details_remarks", "update_time", "is_del").Create(&puchaseDetails).Error
	return err
}
