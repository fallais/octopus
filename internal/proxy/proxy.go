package proxy

import "net"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Proxy is the holder of the configuration.
type Proxy struct {
	AllowedPorts        []int       `json:"allowed_ports" mapstructure:"allowed_ports"`
	AllowedNetworks     []net.IPNet `json:"allowed_networks" mapstructure:"allowed_networks"`
	AllowedMethods      []string    `json:"allowed_methods" mapstructure:"allowed_methods"`
	Whitelist           []string    `json:"whitelist" mapstructure:"whitelist"`
	Blacklist           []string    `json:"blacklist" mapstructure:"blacklist"`
	DisableForwardedFor bool        `json:"disable_forwarded_for" mapstructure:"disable_forwarded_for"`
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewProxy returns a new Proxy.
func NewProxy(whitelist, blacklist []string) (*Proxy, error) {
	// Create the model
	p := &Proxy{
		Blacklist: blacklist,
	}

	return p, nil
}
