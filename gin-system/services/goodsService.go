package services

import (
	"gin-system/dao"
	"gin-system/models"
)

func GetAllGoods(goodsParam map[string]interface{}) (error, []models.Goods, int64) {
	err, goodsData, total := dao.GetAllGoods(goodsParam)
	return err, goodsData, total
}


func DelGoods(id int) error {
	err := dao.DelGoods(id)
	return err
}

func UpdateGoods(goods models.Goods) error {
	err := dao.UpdateGoods(goods)
	return err
}