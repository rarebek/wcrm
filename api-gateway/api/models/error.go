package models

// Error ...
type Error struct {
	Message string `json:"message"`
}

// StandartError ...
type StandartError struct {
	Error Error `json:"error"`
}

type ResponseError struct {
	Code  string      `json:"code"`
	Error interface{} `json:"error"`
}