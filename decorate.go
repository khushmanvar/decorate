package decorate

// Generic Decorate function
func Decorate[T any](base T, decorators ...func(T) T) T {
	decorated := base
	for _, decorator := range decorators {
		decorated = decorator(decorated)
	}
	return decorated
}
