package proxy

import (
	"net/http"
	"os"
	"time"

	"octopus/internal/cache"

	"github.com/spf13/viper"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Proxy is the holder of the configuration.
type Proxy struct {
	opts         options
	cacheManager *CacheManager
	httpClient   *http.Client
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewProxy returns a new Proxy.
func NewProxy(c cache.Cache, opts ...Option) (*Proxy, error) {
	// Create the CacheManager
	cm := &CacheManager{
		IsEnabled: true,
		Cache:     c,
	}

	// Create the HTTP client
	httpClient := &http.Client{
		Timeout:   time.Second * 5,
		Transport: cm,
	}

	// Create the proxy
	p := &Proxy{
		cacheManager: cm,
		httpClient:   httpClient,
	}

	// Set options
	for _, opt := range opts {
		opt(&p.opts)
	}

	// Set the hostname
	hostname, err := os.Hostname()
	if err != nil {
		p.opts.visibleHostname = viper.GetString("general.visible_hostname")
	} else {
		p.opts.visibleHostname = hostname
	}

	return p, nil
}
