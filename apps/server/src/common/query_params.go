package common

// ListQueryParams provides options for retrieving lists of data.
type ListQueryParams struct {
	Page      int    `form:"page" example:"1" binding:"gte=1" default:"1"`
	Limit     int    `form:"limit" example:"25" binding:"gte=1,lte=100" default:"10"`
	SortField string `form:"sort_field" example:"created_at" default:"created_at"`
	SortOrder string `form:"sort_order" example:"desc" binding:"oneof=asc desc" default:"desc"`
	Search    string `form:"search" example:"coffee"`
	Preload   string `form:"preload" example:"category,account"`
}
