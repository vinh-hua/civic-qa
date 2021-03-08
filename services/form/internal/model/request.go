package model

// NewFormRequest is a model for
// a request to create a new form
type NewFormRequest struct {
	Name string `json:"name"`
}

// NewFormResponse is a model for
// a request to create a new response
type NewFormResponse struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// PatchResponse is a model for
// a request to patch a response state
type PatchResponse struct {
	Active bool `json:"active"`
}
