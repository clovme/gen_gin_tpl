package with

import (
	"fmt"
	"gen_gin_tpl/pkg/constants"
	"gen_gin_tpl/pkg/enums/code"
	"gen_gin_tpl/pkg/variable"
	"net/http"
)

// response 响应结构体
type response struct {
	Code    code.ResponseCode `json:"code"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data"`
}

// setResponse 设置响应头
func (r *Context[T]) setHeaderEncryptedResponse(isEnableEncrypted bool) {
	r.isContextEncrypted = isEnableEncrypted
	if !isEnableEncrypted || !variable.IsEnableEncryptedKey {
		r.Context.Header(constants.HeaderEncrypted, "no")
	} else {
		r.Context.Header(constants.HeaderEncrypted, "safe")
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
func (r *Context[T]) JsonSafe(httpCode code.ResponseCode, message string, data interface{}) {
	r.setHeaderEncryptedResponse(true)
	r.Context.JSON(http.StatusOK, response{Code: httpCode, Message: fmt.Sprintf("[%d] %s", httpCode.Int(), message), Data: data})
}

// JsonSafeDesc 安全响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - data: 响应数据
//
// 返回值：
//   - 无
func (r *Context[T]) JsonSafeDesc(httpCode code.ResponseCode, data interface{}) {
	r.JsonSafe(httpCode, httpCode.Desc(), data)
}

// JsonSafeSuccess 安全响应成功
// 参数：
//   - c: gin.Context
//   - data: 响应数据
//
// 返回值：
//   - 无
func (r *Context[T]) JsonSafeSuccess(data interface{}) {
	r.JsonSafe(code.Success, code.Success.Desc(), data)
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
func (r *Context[T]) JsonUnSafe(httpCode code.ResponseCode, message string, data interface{}) {
	r.setHeaderEncryptedResponse(false)
	r.Context.JSON(http.StatusOK, response{Code: httpCode, Message: fmt.Sprintf("[%d] %s", httpCode.Int(), message), Data: data})
}

// JsonUnSafeDesc 不安全响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - data: 响应数据
//
// 返回值：
//   - 无
func (r *Context[T]) JsonUnSafeDesc(httpCode code.ResponseCode, data interface{}) {
	r.JsonUnSafe(httpCode, httpCode.Desc(), data)
}

// JsonUnSafeSuccess 不安全响应成功
// 参数：
//   - c: gin.Context
//   - data: 响应数据
//
// 返回值：
//   - 无
func (r *Context[T]) JsonUnSafeSuccess(data interface{}) {
	r.JsonUnSafe(code.Success, code.Success.Desc(), data)
}
