package utils

const (
	CODE_OK          = "OK"
	CODE_BAD_REQUEST = "ERROR"
)
const (
	InternalMessage = "No se pudo completar la operacion, inténtelo más tarde."
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewOkResponse(data interface{}) *Response {
	return &Response{
		Code: CODE_OK,
		Data: data,
	}
}
func NewBadResponse(message string) *Response {
	return &Response{
		Code:    CODE_BAD_REQUEST,
		Message: message,
	}
}
