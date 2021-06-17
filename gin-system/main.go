package main

import (
	"gin-system/dao"
	"gin-system/pkg/setting"
	"gin-system/routers"
)

func init() {
	setting.Setup()
	dao.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	routersInit.Run()
}
