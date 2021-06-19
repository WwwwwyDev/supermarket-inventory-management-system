package services

import (
	"gin-system/dao"
	"gin-system/models"
)

func GetAllPurchaseList(purchaseListParam map[string]interface{}) (error, []models.PurchaseList, int64) {
	err, purchaseListData, total := dao.GetAllPurchaseList(purchaseListParam)
	return err, purchaseListData, total
}

func DelPurchaseList(id int) error {
	err := dao.DelPurchaseList(id)
	return err
}

func UpdatePurchaseList(purchaseList models.PurchaseList) error {
	err := dao.UpdatePurchaseList(purchaseList)
	return err
}

func AddPurchaseList(purchaseList models.PurchaseList) (error) {
	err := dao.AddPurchaseList(purchaseList)
	return err
}


