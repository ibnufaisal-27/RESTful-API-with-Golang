package models

type ResponseData struct {
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Message string      `json:"message,omitempty"`
}