package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Resp 统一响应结构：{ "code": 0, "message": "ok", "data": ... }
type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// OK 返回成功响应。
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Resp{Code: 0, Message: "ok", Data: data})
}

// Fail 返回业务错误响应，并终止后续 handler 链。
// 权限/限流/参数等拒绝路径调用后，必须阻止后续 handler 继续写响应，
// 否则会拼出既有错误体又有成功体。调用方无需再手动 c.Abort()。
func Fail(c *gin.Context, httpStatus, code int, message string) {
	c.JSON(httpStatus, Resp{Code: code, Message: message})
	c.Abort()
}

// PageResult 分页返回结构。
type PageResult struct {
	List     any   `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
}
