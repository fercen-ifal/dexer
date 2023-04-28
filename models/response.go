package models

type Response struct {
	Message string `json:"message"`
}

// Implementação original em: https://github.com/fercen-ifal/fercen/blob/main/errors/index.ts

type ErrorResponse struct {
	Message   string `json:"message"`
	Action    string `json:"action"`
	ErrorCode string `json:"errorCode"`
}
