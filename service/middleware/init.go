package middleware

import "GoHttpServerBestPractice/global"

/*
   功能说明: 添加全局中间件
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 15:42
*/

func InitGlobalMiddleware() {
	middleware := []global.Middleware{
		AccessLog,
		Cors,
	}
	global.Mux.Use(middleware...)
}
