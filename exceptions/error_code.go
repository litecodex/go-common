package exceptions

import (
	"fmt"
	"runtime"
	"strings"
)

type ErrorCode struct {
	Code    int
	Message string
}

// CustomError 定义自定义错误类型
type CustomError struct {
	Code    int
	Message string
	Stack   string
}

// 实现 error 接口
func (e *CustomError) Error() string {
	return fmt.Sprintf("%s\nStack trace:\n%s", e.Message, e.Stack)
}

const DefaultCode int = 0

// Of 创建一个新的自定义错误
func NewErrorCode(code int, message string) *ErrorCode {
	return &ErrorCode{
		Code:    code,
		Message: message,
	}
}

func OfCode(errCode *ErrorCode) *CustomError {
	return &CustomError{
		Code:    errCode.Code,
		Message: errCode.Message,
		Stack:   captureStackTrace(),
	}
}

func OfError(err error) *CustomError {
	return &CustomError{
		Code:    DefaultCode,
		Message: err.Error(),
		Stack:   captureStackTrace(),
	}
}

// Of 创建一个新的自定义错误
func Of(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Stack:   captureStackTrace(),
	}
}

func OfMessage(message string) *CustomError {
	return &CustomError{
		Code:    DefaultCode,
		Message: message,
		Stack:   captureStackTrace(),
	}
}

func captureStackTrace() string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])

	count := 0

	var sb strings.Builder
	for _, pc := range pcs[:n] {
		count++
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		sb.WriteString(fmt.Sprintf("%s:%d\n", file, line))
		if count > 10 {
			// 堆栈信息只收集10行，多了没必要
			break
		}
	}
	return sb.String()
}
