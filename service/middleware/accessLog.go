package middleware

import (
	"GoHttpServerBestPractice/global"
	"net/http"
	"time"
)

/*
   功能说明: 访问控制
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/18 14:09
*/

func AccessLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//global.AccessLog.Printf("访问日志")
		defer func(start time.Time) {
			global.AccessLog.Printf("%v, %v, %v, %v ms", r.Method, r.URL.RequestURI(), r.RemoteAddr, time.Since(start).Milliseconds())
		}(time.Now())
		handler.ServeHTTP(w, r)
	})
}
