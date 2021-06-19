package models

import "time"

type Goods struct {
	GoodsId           int       `gorm:"column:goods_id"`
	GoodsName         string    `gorm:"column:goods_name"`
	GoodsPrice        int       `gorm:"column:goods_price"`
	GoodsSupplier     int       `gorm:"column:goods_supplier"`
	GoodsSupplierName string    `gorm:"column:goods_supplier_name"`
	GoodsSynopsis     string    `gorm:"column:goods_synopsis"`
	GoodsRemarks      string    `gorm:"column:goods_remarks"`
	Isdel            bool      `gorm:"column:is_del"`
	UpdateTime       time.Time `gorm:"column:update_time"`
}
