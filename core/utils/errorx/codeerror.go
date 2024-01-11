package errorx

import (
	"errors"
	"fmt"
	"net/http"
)

// custom http code
const (
	authCaptcha  = 10003
	tokenInvalid = 40100
)

type (
	CodeError interface {
		error
		Status() int
		Code() int
	}

	codeError struct {
		code   int
		status int
		desc   string
	}
)

func (err *codeError) Code() int {
	return err.code
}

func (err *codeError) Status() int {
	return err.status
}

func (err *codeError) Error() string {
	return err.desc
}

func (err *codeError) String() string {
	return fmt.Sprintf("Status: %d, Code: %d, Desc: %s", err.status, err.code, err.desc)
}

func NewCodeError(code, status int, desc string) CodeError {
	return &codeError{
		code:   code,
		status: status,
		desc:   desc,
	}
}

func FromError(err error) (CodeError, bool) {
	if err == nil {
		return nil, false
	}
	var ce CodeError
	if ok := errors.As(err, &ce); ok {
		return ce, ok
	}
	return nil, false
}

func Error400(desc string) CodeError {
	return NewCodeError(http.StatusBadRequest, http.StatusBadRequest, desc)
}

func Error400f(format string, a ...any) CodeError {
	return NewCodeError(http.StatusBadRequest, http.StatusBadRequest, fmt.Sprintf(format, a))
}

func Error500(desc string) CodeError {
	return NewCodeError(500, 500, desc)
}

func Error500f(format string, a ...any) CodeError {
	return NewCodeError(500, 500, fmt.Sprintf(format, a))
}

func ErrorDB(err error) CodeError {
	return NewCodeError(500, 500, fmt.Sprintf("[DB ERROR]: %v", err))
}

func ErrorRedis(err error) CodeError {
	return NewCodeError(500, 500, fmt.Sprintf("[REDIS ERROR]: %v", err))
}

func ErrorAuth() CodeError {
	return NewCodeError(401, 401, "认证失败")
}

func ErrorTokenInvalid() CodeError {
	return NewCodeError(tokenInvalid, 401, "非法Token")
}

func ErrorAuthCaptcha() CodeError {
	return NewCodeError(authCaptcha, 401, "验证码错误")
}
