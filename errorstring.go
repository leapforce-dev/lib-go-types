package types

// ErrorString : type allowing custom error string
type ErrorString struct {
	Message string
}

// Error: return error message
func (e *ErrorString) Error() string {
	return e.Message
}
