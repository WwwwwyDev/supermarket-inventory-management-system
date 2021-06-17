package main

import (
	"fmt"
	"gin-system/dao"
	"gin-system/pkg/setting"
	"gin-system/routers"
)

func init() {
	setting.Setup()
	dao.Setup()
}

func main() {

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	routersInit := routers.InitRouter()
	routersInit.Run(endPoint)
}
