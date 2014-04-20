package main

import (
	"encoding/xml"
	"fmt"
)

const (
	ErrCodeNotExist      = 1
	ErrCodeAlreadyExists = 2
)

// the serializable Error structure
type Error struct {
	XMLName xml.Name `json:"_" xml:"error"`
	Code    int      `json:"code" xml:"code,attr"`
	Message string   `json:"message" xml:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// create an error instance with specified code and message
func NewError(code int, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
