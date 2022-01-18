package base

import (
	"GoHttpServerBestPractice/global"
	"GoHttpServerBestPractice/service/api/base/controller"
	"GoHttpServerBestPractice/service/middleware"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 14:08
*/

func Router() {
	global.Mux.HandleFunc("/health", controller.HealthHandler)              // 健康检查
	global.Mux.HandleFunc("/test", middleware.Part(controller.TestHandler)) // 局部中间件测试
}
