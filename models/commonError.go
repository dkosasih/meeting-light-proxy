package models

// swagger:response CommonError
type CommonError struct {
	// success status
	// in: bool
	Success bool `json:"success"`
	// error message
	// in: string
	Error string `json:"error"`
}
