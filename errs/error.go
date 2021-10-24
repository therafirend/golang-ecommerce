package errs

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (a AppError) GetMsg() *AppError {
	return &AppError{
		Message: a.Message,
	}
}

func NewAppError(msg string, code int) *AppError {
	return &AppError{
		Message: msg,
		Code:    code,
	}
}
