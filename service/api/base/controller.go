package base

import (
	"GoHttpServerBestPractice/service/core"
	"net/http"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 14:09
*/

func healthHandler(w http.ResponseWriter, r *http.Request) {
	core.Handler200(w, nil)
	return
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	core.Handler200(w, "test")
	return
}
