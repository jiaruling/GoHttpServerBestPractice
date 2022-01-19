package user

import (
	"GoHttpServerBestPractice/global"
	"GoHttpServerBestPractice/service/api/user/controller"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 14:06
*/

func Router() {
	global.Mux.HandleFunc("/stu", controller.StuHandler)
}
