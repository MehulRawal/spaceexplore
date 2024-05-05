package api

type Response struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// Message may contain a string that we require to show on the toast in UI
func ResponseConstruct(message string, errorMessage string, success bool, data interface{}) Response {
	return Response{Message: message, Error: errorMessage, Success: success, Data: data}
}
