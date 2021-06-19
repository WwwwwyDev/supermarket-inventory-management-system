package services

import (
	"gin-system/dao"
	"gin-system/models"
)

func GetAllPurchaseDetails(purchaseDetails map[string]interface{}) (error, []models.PuchaseDetails, int64) {
	err, purchaseDetailsData, total := dao.GetAllPurchaseDetails(purchaseDetails)
	return err, purchaseDetailsData, total
}

func DelPurchaseDetails(id int) error {
	err := dao.DelPurchaseDetails(id)
	return err
}

func UpdatePurchaseDetails(puchaseDetails models.PuchaseDetails) error {
	err := dao.UpdatePurchaseDetails(puchaseDetails)
	return err
}

func AddPurchaseDetails(puchaseDetails models.PuchaseDetails) (error) {
	err := dao.AddPurchaseDetails(puchaseDetails)
	return err
}
