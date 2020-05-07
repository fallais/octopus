package proxy

import (
	"net"
	"time"

	"octopus/internal/cache"
	"octopus/internal/cache/local"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Cache is the cache configuration.
type Cache struct {
	IsEnabled bool
	Type      string
}

// Proxy is the holder of the configuration.
type Proxy struct {
	AllowedPorts        []int
	AllowedNetworks     []net.IPNet
	AllowedMethods      []string
	Whitelist           []string
	Blacklist           []string
	DisableForwardedFor bool
	Cache               cache.Cache
	IsCacheEnabled      bool
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewProxy returns a new Proxy.
func NewProxy(whitelist, blacklist []string) (*Proxy, error) {
	// Cache
	c, _ := local.NewLocalCache("configs\\cache\\", 1*time.Hour)

	// Create the model
	p := &Proxy{
		Blacklist:      blacklist,
		IsCacheEnabled: true,
		Cache:          c,
	}

	return p, nil
}
