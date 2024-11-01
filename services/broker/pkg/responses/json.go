package responses

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func ErrorResponse(err error) JSONResponse {
	return JSONResponse{
		Error:   true,
		Message: err.Error(),
	}
}
