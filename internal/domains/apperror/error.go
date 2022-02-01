package apperror

import "encoding/json"

type AppError struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

var (
	InternalError *AppError = NewAppError("message", "S-000001")
)

func NewAppError(message, code string) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func (e *AppError) Error() string { return e.Message }

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}
