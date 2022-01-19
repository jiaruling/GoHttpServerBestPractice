package controller

import (
	"GoHttpServerBestPractice/service/api/user/model"
	"GoHttpServerBestPractice/service/core"
	"net/http"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 14:06
*/

func StuHandler(w http.ResponseWriter, r *http.Request) {
	// 通过全局变量赋值给局部变量，可以实现并发
	s := model.Stu
	s.M = new(model.Student)
	core.Dispatcher(s, w, r)
	return
}
