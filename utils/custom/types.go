package custom

// JSONResponse is a template for all the api responses
type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// RequestError implements error interface for request related errors
type RequestError struct {
	Status  int // https status code
	Message string
}

func (err RequestError) Error() string {
	return err.Message
}
