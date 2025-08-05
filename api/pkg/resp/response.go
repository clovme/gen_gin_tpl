package resp

import (
	"fmt"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/variable"
	"github.com/gin-gonic/gin"
	"net/http"
)

// response 响应结构体
type response struct {
	Code    code.ResponseCode `json:"code"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data"`
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

// JsonSafe 安全响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - message: 响应消息
//   - data: 响应数据
//
// 返回值：
//   - 无
func JsonSafe(c *gin.Context, httpCode code.ResponseCode, message string, data interface{}) {
	setResponse(c, true)
	c.JSON(http.StatusOK, response{Code: httpCode, Message: fmt.Sprintf("[%d] %s", httpCode.Enum(), message), Data: data})
}

// JsonSafeDesc 安全响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - data: 响应数据
//
// 返回值：
//   - 无
func JsonSafeDesc(c *gin.Context, httpCode code.ResponseCode, data interface{}) {
	JsonSafe(c, httpCode, httpCode.Desc(), data)
}

// JsonSafeSuccess 安全响应成功
// 参数：
//   - c: gin.Context
//   - data: 响应数据
//
// 返回值：
//   - 无
func JsonSafeSuccess(c *gin.Context, data interface{}) {
	JsonSafe(c, code.Success, code.Success.Desc(), data)
}

// JsonUnSafe 不安全响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - message: 响应消息
//   - data: 响应数据
//
// 返回值：
//   - 无
func JsonUnSafe(c *gin.Context, httpCode code.ResponseCode, message string, data interface{}) {
	setResponse(c, false)
	c.JSON(http.StatusOK, response{Code: httpCode, Message: fmt.Sprintf("[%d] %s", httpCode.Enum(), message), Data: data})
}

// JsonUnSafeDesc 不安全响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - data: 响应数据
//
// 返回值：
//   - 无
func JsonUnSafeDesc(c *gin.Context, httpCode code.ResponseCode, data interface{}) {
	JsonUnSafe(c, httpCode, httpCode.Desc(), data)
}

// JsonUnSafeSuccess 不安全响应成功
// 参数：
//   - c: gin.Context
//   - data: 响应数据
//
// 返回值：
//   - 无
func JsonUnSafeSuccess(c *gin.Context, data interface{}) {
	JsonUnSafe(c, code.Success, code.Success.Desc(), data)
}
