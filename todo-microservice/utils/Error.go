package utils

// ApiErrors
// Response structure when there will be any api errors
type ApiErrors struct {
	Errors     []FieldError `json:"errors"`
	Message    string       `json:"message"`
	StatusCode int          `json:"status_code"`
}

// FieldError
// It will be thrown as an Array in ApiError
type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
