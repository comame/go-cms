package src

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"error"`
	Code    string `json:"code"`
}

type ErrorCode int

const (
	NotFound ErrorCode = iota
	InternalError
)

func (self ErrorCode) String() string {
	switch self {
	case NotFound:
		return "NotFound"
	case InternalError:
		return "InternalError"
	default:
		panic("unreachable")
	}
}
