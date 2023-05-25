package errorcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code_    int      `json:"code"`
	Msg_     string   `json:"msg"`
	Details_ []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("錯誤%d已經存在，請更換一個", code))
	}
	codes[code] = msg

	return &Error{Code_: code, Msg_: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("錯誤:%d,錯誤訊息:%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.Code_
}

func (e *Error) Msg() string {
	return e.Msg_
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg_, args...)
}

func (e *Error) Details() []string {
	return e.Details_
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details_ = []string{}

	newError.Details_ = append(newError.Details_, details...)

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code_ {
	case Success.Code_:
		return http.StatusOK
	case ServerError.Code_:
		return http.StatusInternalServerError
	case InvalidParams.Code_:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code_:
		fallthrough
	case UnauthorizedTokenError.Code_:
		fallthrough
	case UnauthorizedTokenGenerate.Code_:
		fallthrough
	case UnauthorizedTokenTimeout.Code_:
		return http.StatusUnauthorized
	case TooManyRequests.Code_:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
