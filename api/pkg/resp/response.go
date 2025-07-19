package resp

import (
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/enums/em_http"
	"gen_gin_tpl/pkg/variable"
	"github.com/gin-gonic/gin"
	"net/http"
)

// response 响应结构体
type response struct {
	Code    em_http.ResponseCode `json:"code"`
	Message string               `json:"message"`
	Data    interface{}          `json:"data"`
}

// setResponse 设置响应头
func setResponse(c *gin.Context, flag bool) {
	c.Set(constants.ContextIsEncrypted, flag)
	if !flag || !variable.IsEnableEncryptedKey {
		c.Header(constants.HeaderEncrypted, "no")
	} else {
		c.Header(constants.HeaderEncrypted, "safe")
	}
}

// JsonSafeCode 安全响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - message: 响应消息
//   - obj: 响应数据
//
// 返回值：
//   - 无
func JsonSafeCode(c *gin.Context, code em_http.ResponseCode, message string, obj interface{}) {
	setResponse(c, true)
	c.JSON(http.StatusOK, response{Code: code, Message: message, Data: obj})
}

// JsonSafe 安全响应
// 参数：
//   - c: gin.Context
//   - message: 响应消息
//   - obj: 响应数据
//
// 返回值：
//   - 无
func JsonSafe(c *gin.Context, message string, obj interface{}) {
	JsonSafeCode(c, em_http.Success, message, obj)
}

// JsonUnSafeCode 不安全响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - message: 响应消息
//   - obj: 响应数据
//
// 返回值：
//   - 无
func JsonUnSafeCode(c *gin.Context, code em_http.ResponseCode, message string, obj interface{}) {
	setResponse(c, false)
	c.JSON(http.StatusOK, response{Code: code, Message: message, Data: obj})
}

// JsonUnSafe 不安全响应
// 参数：
//   - c: gin.Context
//   - message: 响应消息
//   - obj: 响应数据
//
// 返回值：
//   - 无
func JsonUnSafe(c *gin.Context, message string, obj interface{}) {
	JsonUnSafeCode(c, em_http.Success, message, obj)
}
