package utils

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	Status  int
	Message string
	Detail  error
}

func (e HttpError) Error() string {
	return fmt.Sprintf("%v", e.Detail)
}

func (e HttpError) GetMessage() string {
	if e.Message == "" {
		return "Internal server error"
	}
	return e.Message
}

func (e HttpError) GetDetail() error {
	return e.Detail
}

func (e HttpError) GetStatus() int {
	if e.Status == 0 {
		return http.StatusInternalServerError
	}
	return e.Status
}
