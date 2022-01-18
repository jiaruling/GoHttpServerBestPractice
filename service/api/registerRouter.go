package api

import (
	"GoHttpServerBestPractice/service/api/base"
	"GoHttpServerBestPractice/service/api/user"
)

/*
   功能说明: 注册路由
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 11:48
*/

func RegisterRouter() {
	base.Router()
	user.Router()
}
