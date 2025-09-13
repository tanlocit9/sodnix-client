package response

import (
	"net/http"
)

func SuccessResponse[T any](data T, code int, message any) GenericResponse[T] {
	if message, ok := message.(string); ok {
		return GenericResponse[T]{Data: &data, Code: code, Message: message}
	} else {
		return GenericResponse[T]{Data: &data, Code: code, Message: http.StatusText(code)}
	}
}

func GetDataSuccessResponse[T any](data T) GenericResponse[T] {
	statusCode := http.StatusOK
	return SuccessResponse(data, statusCode, http.StatusText(statusCode))
}

func CreateDataSuccessResponse[T any](data T) GenericResponse[T] {
	statusCode := http.StatusCreated
	return SuccessResponse(data, statusCode, http.StatusText(statusCode))
}

func UpdateDataSuccessResponse[T any](data T) GenericResponse[T] {
	statusCode := http.StatusOK
	return SuccessResponse(data, statusCode, http.StatusText(statusCode))
}

func DeleteDataSuccessResponse[T any](data T) GenericResponse[T] {
	statusCode := http.StatusNoContent
	return SuccessResponse(data, statusCode, http.StatusText(statusCode))
}
