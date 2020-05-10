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
	ticker            *time.Ticker
	defaultExpiration time.Duration
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewLocalCache returns a new local cache.
func NewLocalCache(path string, defaultExpiration time.Duration) (cache.Cache, error) {
	lc := &localCache{
		path:              path,
		ticker:            time.NewTicker(1 * time.Second),
		defaultExpiration: defaultExpiration,
	}

	go lc.watch()

	return lc, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

func (c *localCache) watch() {
	for range c.ticker.C {
		go c.clean()
	}
}

func (c *localCache) clean() {

}

// Set a value in cache.
func (c *localCache) Set(key string, value []byte) error {
	return ioutil.WriteFile(c.path+"\\"+key, value, 0644)
}

// Get a value from cache.
func (c *localCache) Get(key string) ([]byte, error) {
	return ioutil.ReadFile(c.path + "\\" + key)
}
