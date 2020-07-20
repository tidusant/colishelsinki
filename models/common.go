package models

type RequestResult struct {
	Status  string          `json:"status"`
	Error   string          `json:"error"`
	Message string          `json:"message"`
	Data    string `json:"data"`
}
