package utilities

// PtrTo takes in a value and returns a pointer address to that value. This is largely used in order to concisely
// initialize pointers from constant primitives in struct initialization. This should probably only be used with
// primitives, as more complex types can just use "&". Note: this function does NOT work with float32.
func PtrTo[T any](s T) *T {
	return &s
}
