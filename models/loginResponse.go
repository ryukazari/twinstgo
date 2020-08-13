package models

// LoginResponse struct for login token
type LoginResponse struct {
	Ok      bool   `json:"ok"`
	Status  int    `json:"status"`
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
}
