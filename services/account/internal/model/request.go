package model

// NewUserRequest represents a json POST
// request to /signup
type NewUserRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
}

// LoginRequest represents a json POST
// request to /login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
