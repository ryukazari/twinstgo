package models

// Response struct for default response
type Response struct {
	Ok      bool        `json:"ok"`
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
