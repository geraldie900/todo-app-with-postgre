package utils

// Response parent object for response
// use this as all purpose response object
type Response struct {
	ResponseData interface{}
	StatusCode   int
}

// JsonResponse json object for response
type JsonResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// GenerateJsonResponse used to generate json response
func GenerateJsonResponse(success bool, message string, data interface{}, err interface{}) JsonResponse {
	response := JsonResponse{}

	response.Success = success
	response.Message = message
	response.Data = data
	response.Error = err

	return response
}
