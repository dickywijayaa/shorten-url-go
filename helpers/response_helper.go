package helpers

import (
	"../objects"

	"net/http"
)

type ResponseHelper struct {
}

func ResponseHelperHandler() ResponseHelper {
	return ResponseHelper{}
}

func (h *ResponseHelper) BadRequestResponse(data interface{}, message string) objects.Response {
	response := objects.Response{
		Code: http.StatusBadRequest,
		Data: data,
		Message: message,
		Status: "Failed",
	}

	return response
}

func (h *ResponseHelper) SuccessResponse(data interface{}, message string) objects.Response {
	response := objects.Response{
		Code: http.StatusOK,
		Data: data,
		Message: message,
		Status: "Success",
	}

	return response
}