package dto

type UserRequest struct {
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Country string `json:"country,omitempty"`
}

type UserResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}
