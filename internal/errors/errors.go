package errors

import (
	"fmt"
	"net/http"
)

// Error 自定义错误类型
type Error struct {
	Code    int    // HTTP 状态码
	Message string // 错误信息
	Err     error  // 原始错误
}

// Error 实现 error 接口
func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap 返回原始错误
func (e *Error) Unwrap() error {
	return e.Err
}

// New 创建新的错误
func New(code int, message string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// IsNotFound 检查是否为"未找到"错误
func IsNotFound(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == http.StatusNotFound
	}
	return false
}

// IsBadRequest 检查是否为"错误请求"错误
func IsBadRequest(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == http.StatusBadRequest
	}
	return false
}

// IsInternalServer 检查是否为"服务器内部错误"
func IsInternalServer(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.Code == http.StatusInternalServerError
	}
	return false
}

// Wrap 包装错误
func Wrap(err error, message string) *Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		return New(e.Code, message, e)
	}
	return New(http.StatusInternalServerError, message, err)
} 