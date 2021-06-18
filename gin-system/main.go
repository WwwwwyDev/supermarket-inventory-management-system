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

	//for i:=1; i < 10000; i++{
	//	services.AddStaff(models.Staff{1,"testtest" + strconv.Itoa(i),"wwy123456",1,"2312312",132133,"测试用户",false,time.Now()})
	//}
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	routersInit := routers.InitRouter()
	routersInit.Run(endPoint)
}
