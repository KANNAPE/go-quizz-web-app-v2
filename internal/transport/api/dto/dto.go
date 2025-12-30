package dto

import (
	"net/http"
	"time"
)

type APIResponse[T any] struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      T         `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func NewAPIResponse[T any]() APIResponse[T] {
	return APIResponse[T]{
		Code:      http.StatusOK,
		Message:   "success",
		Timestamp: time.Now(),
	}
}
