package converter

// Generic function to concatenate multiple slices of type T
func ConcatAll[T any](slices ...[]T) []T {
	var result []T
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

// Converts a []interface{} to []any (mostly for Go 1.18+ compatibility)
func ToAny[T any](slice []T) []any {
	result := make([]any, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

// ConcatToAny takes multiple slices of any type and returns a single []any
func ConcatToAny[T any](slices ...[]T) []any {
	return ToAny(ConcatAll(slices...))
}

// ConcatAny takes multiple slices of []any and returns a single []any
func ConcatAny(slices ...[]any) []any {
	var result []any
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}
