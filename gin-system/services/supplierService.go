package services

import (
	"gin-system/dao"
	"gin-system/models"
)

func GetAllSupplier(supplierParam map[string]interface{}) (error, []models.Supplier, int64) {
	err, supplierData, total := dao.GetAllSupplier(supplierParam)
	return err, supplierData, total
}

func DelSupplier(id int) error {
	err := dao.DelSupplier(id)
	return err
}

func UpdateSupplier(supplier models.Supplier) (error) {
	err := dao.UpdateSupplier(supplier)
	return err
}

func AddSupplier(supplier models.Supplier) (error) {
	err := dao.AddSupplier(supplier)
	return err
}