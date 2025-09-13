package common

// GetOptions provides options for retrieving a single record.
type GetOptions struct {
	Preload string `form:"preload"`
}
