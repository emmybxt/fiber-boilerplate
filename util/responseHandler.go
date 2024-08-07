package util


func SuccessResponse(message string, data interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["message"] = message
	response["data"] = data
	response["success"] = true
	return response
}


func ErrorResponse(message string) map[string]interface{} {
	response := make(map[string]interface{})
	response["message"] = message
	response["success"] = false

	return response
}