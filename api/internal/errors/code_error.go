package errors

const defaultCode = 3

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewCodeWrongPassError(msg string) error {
	return &CodeError{Code: 1, Msg: msg}
}

func NewCodeWrongCaptcha(msg string) error {
	return &CodeError{
		Code: 2,
		Msg:  msg,
	}
}

// 自定义的错误返回
func NewCodeAbortedError(msg string) error {
	return &CodeError{Code: 10, Msg: msg}
}

func (e *CodeError) Error() string {
	return e.Msg
}
