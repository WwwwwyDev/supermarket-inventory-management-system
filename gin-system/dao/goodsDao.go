package dao

import (
	"gin-system/models"
)

func GetAllGoods(goodsParam map[string]interface{}) (error, []models.Goods, int64) {
	var goodsData []models.Goods
	page := goodsParam["page"].(int)
	pageSize := goodsParam["limit"].(int)
	searchName := goodsParam["searchName"].(string)
	var total int64
	err := db.Table("erp_goods g").Select("g.goods_id, g.goods_name, g.goods_price, g.goods_supplier,s.supplier_name goods_supplier_name, g.goods_remarks,g.goods_synopsis,g.is_del, g.update_time").Joins("INNER JOIN erp_supplier s ON g.goods_supplier = s.supplier_id").Where("g.goods_name like ? and g.is_del = false", "%"+searchName+"%").Order("goods_id ASC").Count(&total).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsData).Error
	return err, goodsData, total
}

func DelGoods(id int) error {
	err := db.Table("erp_goods").Where("goods_id = ?", id).Update("is_del", true).Error
	return err
}

func UpdateGoods(goods models.Goods) error {
	err := db.Table("erp_goods").Where("goods_id = ?", goods.GoodsId).Updates(&goods).Error
	return err
}

func AddGoods(goods models.Goods) (error) {
	err := db.Table("erp_goods").Select("goods_name", "goods_price", "goods_supplier", "goods_synopsis", "goods_remarks", "is_del", "update_time").Create(&goods).Error
	return err
}
