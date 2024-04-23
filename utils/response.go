package utils

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func StructuredResponse(success bool, message string, data interface{}) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func SuccessResponse(message string, data interface{}) Response {
	return StructuredResponse(true, message, data)
}

func ErrorResponse(message string, data interface{}) Response {
	return StructuredResponse(false, message, data)
}
