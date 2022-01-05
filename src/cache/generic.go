package cache
type Cache[T any] struct {
    data map[string]T
}
func New[T any]() *Cache[T] {
    return &Cache[T]{data: make(map[string]T)}
}