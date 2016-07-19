package errors

const (
	BADREQUEST_CODE = 400
	UNAUTHORIZE_CODE = 401
	FORBIDDEN_CODE = 403
	NOTFOUND_CODE = 404
)

type Errors struct {
	Status  int
	Message string
}

func (e *Errors) Error() string {
	return e.Message
}

func NewErrors(message string) error {
	return &Errors{Message: message}
}

func BadRequestError(message string) error {
	return &Errors{Status: BADREQUEST_CODE, Message: message}
}

func UnAuthorizeError(message string) error {
	return &Errors{Status: UNAUTHORIZE_CODE, Message: message}
}

func ForbiddenError(message string) error {
	return &Errors{Status: FORBIDDEN_CODE, Message: message}
}

func NotFoundError(message string) error {
	return &Errors{Status: NOTFOUND_CODE, Message: message}
}