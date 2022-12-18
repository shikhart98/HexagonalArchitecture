package errs

import "net/http"

type AppErrors struct {
	Code int		`json:"code,omitempty"`
	Message string 	`json:"message"`
}


func (e AppErrors) AsMessage() *AppErrors {
	return &AppErrors{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppErrors {
	return &AppErrors{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) *AppErrors {
	return &AppErrors{
		Code: http.StatusInternalServerError,
		Message: message,
	}
}

func NewCustomError(code int, message string) *AppErrors {
	return &AppErrors{
		Code: code,
		Message: message,
	}
}