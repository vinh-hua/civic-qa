package model

type NewFormRequest struct {
	Name string `json:"name"`
}

type NewFormResponse struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
