package middleware

import (
	"GoHttpServerBestPractice/global"
	"net/http"
)

/*
   功能说明: 局部中间件
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 16:02
*/
func Part(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		global.AccessLog.Println("局部中间件")
		handler.ServeHTTP(w, r)
	}
}
