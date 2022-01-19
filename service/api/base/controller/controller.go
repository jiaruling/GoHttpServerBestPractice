package controller

import (
	"GoHttpServerBestPractice/global"
	"GoHttpServerBestPractice/global/errInfo"
	"GoHttpServerBestPractice/service/core"
	"GoHttpServerBestPractice/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
   功能说明:
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1/18 14:09
*/

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	core.Handler200(w, nil)
	return
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	core.Handler200(w, "test")
	return
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// 判断请求方式
	if r.Method != global.METHOD_POST {
		core.Handler400(w, errInfo.RequestNotAllow, nil)
		return
	}
	// 验证文件大小
	r.Body = http.MaxBytesReader(w, r.Body, global.MaxUploadSize)
	if err := r.ParseMultipartForm(global.MaxUploadSize); err != nil {
		core.Handler400(w, errInfo.FileIsTooBig, nil)
		return
	}

	// 获取文件名和文件内容
	file, fileHandle, err := r.FormFile("file")
	if err != nil {
		core.Handler400(w, errInfo.FileGetFailed, nil)
		return
	}
	defer func() { _ = file.Close() }()
	filePath, fileName := generateFileName(fileHandle.Filename)
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		core.Handler500(w, errInfo.FileReadFailed)
		return
	}
	//fmt.Println(filePath, fileName)
	// 判断文件是否存在，存在就将其删除
	if utils.Exists(filePath) {
		_ = os.Remove(filePath)
	}
	// 将数据写入文件
	newFile, err := os.Create(filePath)
	if err != nil {
		core.Handler500(w, errInfo.FileCreateFailed)
		return
	}
	defer func() { _ = newFile.Close() }()
	if _, err = newFile.Write(fileBytes); err != nil {
		core.Handler500(w, errInfo.FileWriteFailed)
		return
	}
	// 成功返回
	core.Handler201(w, map[string]string{
		"filename": fileName,
	})
	return
}

// 生成文件名
func generateFileName(fileNameOrigin string) (filePath, fileOrigin string) {
	stringList := strings.Split(fileNameOrigin, ".")
	if len(stringList) == 2 {
		fileOrigin = fmt.Sprintf("%s%v.%s", stringList[0], time.Now().Unix(), stringList[1])
	} else {
		fileOrigin = fmt.Sprintf("%s%v", stringList[0], time.Now().Unix())
	}
	filePath = fmt.Sprintf("%s/%s", global.FILEPATH, fileOrigin)
	return
}
