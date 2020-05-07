package local

import (
	"io/ioutil"
	"time"

	"octopus/internal/cache"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type localCache struct {
	path              string
	defaultExpiration time.Duration
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewLocalCache returns a new local cache.
func NewLocalCache(path string, defaultExpiration time.Duration) (cache.Cache, error) {
	return &localCache{
		path:              path,
		defaultExpiration: defaultExpiration,
	}, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Set a value in cache.
func (c *localCache) Set(key string, value interface{}, expires time.Duration) error {
	return ioutil.WriteFile(c.path+"\\"+key, value.([]byte), 0644)
}

// Get a value from cache.
func (c *localCache) Get(key string) ([]byte, error) {
	return ioutil.ReadFile(c.path + "\\" + key)
}
