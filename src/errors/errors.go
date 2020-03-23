package errors

import (
	"reflect"
)

//ErrorCode - error structure
type ErrorCode struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//Recover - recover from panic
func Recover(e ErrorCode) {
	if err := recover(); err != nil {
		val := reflect.ValueOf(err)
		e.Message = val.String()
		e.Code = 500
	}
}
