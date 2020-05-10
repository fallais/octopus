package cache

// Cache ...
type Cache interface {
	Get(string) ([]byte, error)
	Set(string, []byte) error
}
