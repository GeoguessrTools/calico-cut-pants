package utilities

// StrPtr converts a string into a pointer to a string. This allows you to generate a string pointer from a constant
// value without making an intermediary variable first.
func StrPtr(val string) *string {
	return &val
}
