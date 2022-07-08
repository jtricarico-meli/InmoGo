package utils

import (
	"fmt"
)

type InmoError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, message string) error {
	return InmoError{
		Code:    code,
		Message: message,
	}

}

func (e InmoError) Error() string {
	return fmt.Sprintf(`{"code": %v, "message": "%s"}`, e.Code, e.Message)
}

func (e InmoError) ErrorJson() []byte {
	return []byte(fmt.Sprintf(`{"code": %v, "message": "%s"}`, e.Code, e.Message))
}
