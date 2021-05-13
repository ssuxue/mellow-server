package model

import "net/http"

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}
 
func (result *Result) Success(data interface{}) Result {
	return Result{
		Code:    http.StatusOK,
		Message: "操作成功",
		Data:    data,
	}
}
