package models

import "time"

type PuchaseDetails struct {
	PurchaseDetailsId        int       `gorm:"column:purchase_details_id"`
	PurchaseDetailsList      int       `gorm:"column:purchase_details_list"`
	PurchaseDetailsGoods     int       `gorm:"column:purchase_details_goods"`
	PurchaseDetailsGoodsName string    `gorm:"column:purchase_details_goods_name"`
	PurchaseDetailsNumber    int       `gorm:"column:purchase_details_number"`
	PurchaseDetailsPrice     int       `gorm:"column:purchase_details_price"`
	PurchaseDetailsRemarks   string    `gorm:"column:purchase_details_remarks"`
	UpdateTime               time.Time `gorm:"column:update_time"`
	IsDel                    bool      `gorm:"column:is_del"`
}
