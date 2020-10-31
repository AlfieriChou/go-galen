package app

import (
	"github.com/gin-gonic/gin"
)

// Gin gin结构体
type Gin struct {
	C *gin.Context
}

// Result 返回结构体
type Result struct {
	code int      
	message  string      
	data interface{} 
}

// GetMsg 获取错误alias
func getMsg(code string) string {
	var msgFlags map[string]string
	msgFlags = make(map[string]string)
	msgFlags["OK"] = "ok"
	msgFlags["ERROR"] = "failed"
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags["ERROR"]
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, code int, message string, data interface{}) {
	g.C.JSON(httpCode, Result{
		code: code,
		message:  getMsg(message),
		data: data,
	})
	return
}
