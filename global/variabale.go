package global

import (
	. "GoHttpServerBestPractice/global/config_struct"
	"log"
	"net/http"
	"time"
)

/*
   功能说明: 变量
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 10:52
*/

var (
	Config    ServerConfig
	AccessLog *log.Logger
	SqlLog    *log.Logger
	TaskLog   *log.Logger
	LogPath   []string
	Mux       *MyMux
	Server    *http.Server
)

// 初始化全局变量
func init() {
	LogPath = []string{"./log/logs.log", "./log/access.log", "./log/sql.log", "./log/backend_task.log"}
	Mux = NewMyMux()
	Server = &http.Server{
		Addr:         ":8080",
		Handler:      Mux,
		ReadTimeout:  3 * time.Second, // 读超时时间
		WriteTimeout: 3 * time.Second, // 写超时时间
	}
}
