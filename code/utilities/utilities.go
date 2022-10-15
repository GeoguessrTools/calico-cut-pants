package utilities

// StrPtr converts a string into a pointer to a string. This allows you to generate a string pointer from a constant
// value without making an intermediary variable first.
func StrPtr(val string) *string {
	return &val
}

// Float32Ptr converts a float32 into a pointer to a float32. This allows you to generate a float32 pointer from a constant
// value without making an intermediary variable first.
func Float32Ptr(val float32) *float32 {
	return &val
}
