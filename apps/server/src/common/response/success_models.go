package response

type GetDataSuccess[T any] struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Get data success"`
	Data    T      `json:"data,omitempty"`
}

type CreateSuccess[T any] struct {
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"Resource has been created"`
	Data    T      `json:"data,omitempty"`
}

type UpdateSuccess[T any] struct {
	Code    int    `json:"code" example:"204"`
	Message string `json:"message" example:"Resource has been updated"`
	Data    T      `json:"data,omitempty"`
}

type DeleteSuccess[T any] struct {
	Code    int    `json:"code" example:"204"`
	Message string `json:"message" example:"Resource has been deleted"`
	Data    T      `json:"data,omitempty"`
}
