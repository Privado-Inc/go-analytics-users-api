package errors

import "net/http"

type RestErr struct {
	Message string
	Status  int
	Error   string
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "bad_request",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
