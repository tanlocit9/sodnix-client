package response

import (
	"net/http"
)

func ErrResponse[T any](statusCode int, message any) GenericResponse[T] {
	if message, ok := message.(string); ok {
		return GenericResponse[T]{Code: statusCode, Message: message}
	} else {
		return GenericResponse[T]{Code: statusCode, Message: http.StatusText(statusCode)}
	}
}

func NotFoundResponse[T any](message any) GenericResponse[T] {
	return ErrResponse[T](http.StatusNotFound, message)
}

func BadRequestResponse[T any](message any) GenericResponse[T] {
	return ErrResponse[T](http.StatusBadRequest, message)
}

func UnauthorizedResponse[T any](message any) GenericResponse[T] {
	return ErrResponse[T](http.StatusUnauthorized, message)
}

func InternalServerErrorResponse[T any](message any) GenericResponse[T] {
	return ErrResponse[T](http.StatusInternalServerError, message)
}
