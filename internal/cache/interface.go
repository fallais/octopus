package cache

import "time"

// Cache ...
type Cache interface {
	Get(string) ([]byte, error)
	Set(string, interface{}, time.Duration) error
}
