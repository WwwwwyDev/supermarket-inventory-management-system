package dao

import (
	"gin-system/models"
)

func GetAllSupplier(supplierParam map[string]interface{}) (error, []models.Supplier, int64) {
	var supplierData []models.Supplier
	page := supplierParam["page"].(int)
	pageSize := supplierParam["limit"].(int)
	searchName := supplierParam["searchName"].(string)
	var total int64
	err := db.Table("erp_supplier").Where("supplier_name like ? and is_del = false", searchName+"%").Order("supplier_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&supplierData).Error
	return err, supplierData, total
}

func AddSupplier(supplier models.Supplier) error {
	err := db.Table("erp_supplier").Select("supplier_name", "supplier_shortname", "supplier_address", "supplier_ctelephone", "supplier_email", "supplier_liaisonman", "supplier_ltelephone",
		"supplier_remarks", "is_del", "update_time").Create(&supplier).Error
	return err
}

func UpdateSupplier(supplier models.Supplier) error {
	err := db.Table("erp_supplier").Where("supplier_id = ?", supplier.SupplierId).Updates(&supplier).Error
	return err
}
func DelSupplier(id int) error {
	err := db.Table("erp_supplier").Where("supplier_id = ?", id).Update("is_del", true).Error
	return err
}
