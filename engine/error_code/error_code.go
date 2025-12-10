package error_code

type Code struct {
	Code    int
	Message string
}

// 创建错误码的辅函数
func New(code int, msg string) Code {
	return Code{Code: code, Message: msg}
}

func (c Code) Error() string {
	return c.Message
}

var (
	OK              = New(0, "success")
	ErrBadRequest   = New(400, "bad request")
	ErrUnauthorized = New(401, "unauthorized")
	ErrForbidden    = New(403, "forbidden")
	ErrNotFound     = New(404, "not found")
	ErrInternal     = New(500, "internal server error")
	ErrBadMd5       = New(1000, "bad md5 key")
	ErrParam        = New(1001, "ErrParam")
)
