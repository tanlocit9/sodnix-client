package response

// @name BadRequestError
type BadRequestError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

// @name UnauthorizedError
type UnauthorizedError struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"Unauthorized"`
}

// @name NotFoundError
type NotFoundError struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Resource not found"`
}

// @name InternalServerError
type InternalServerError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Internal Server Error"`
}
