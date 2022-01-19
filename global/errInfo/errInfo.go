package errInfo

/*
   功能说明: 错误提示信息
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 13:45
*/

type ErrInfo = string

const (
	RequestNotAllow   ErrInfo = "请求方式不被允许"
	FileIsTooBig      ErrInfo = "文件太大"
	FileGetFailed     ErrInfo = "文件获取失败"
	FileNameIsInvalid ErrInfo = "文件名无效"
	FileReadFailed    ErrInfo = "文件读取失败"
	FileCreateFailed  ErrInfo = "文件创建失败"
	FileWriteFailed   ErrInfo = "文件写入失败"
)
