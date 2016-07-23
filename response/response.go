package response

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func DefaultResponse(message string) *Response {
	return &Response{Message: message}
}

func ShortResponse(message string, data interface{}) *Response {
	return &Response{Message: message, Data: data}
}

func NewResponse(message string, code int, data interface{}) *Response {
	return &Response{Message: message, Code: code, Data: data}
}