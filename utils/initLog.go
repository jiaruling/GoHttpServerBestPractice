package utils

import (
	"GoHttpServerBestPractice/global"
	"fmt"
	"log"
	"os"
)

func DeleteLogFile(LogList []string) {
	for _, item := range LogList {
		if Exists(item) {
			_ = os.Remove(item)
		}
	}
}

// 初始化自带的日志库
func InitLog(LogPath []string) {
	for i, item := range LogPath {
		logFile, err := os.OpenFile(item, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("open log file failed, err:", err)
			return
		}
		//defer logFile.Close()
		switch i {
		case 0:
			log.SetOutput(logFile)
			log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
		case 1:
			global.AccessLog = log.New(logFile, "<Access>: ", log.Lshortfile|log.Ldate|log.Ltime)
		case 2:
			global.SqlLog = log.New(logFile, "<SQL>: ", log.Lshortfile|log.Ldate|log.Ltime)
		}
	}
}
