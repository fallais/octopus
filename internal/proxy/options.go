package proxy

import "net"

// options are the options for the proxy.
type options struct {
	whitelist            []string
	blacklist            []string
	visibleHostname      string
	allowedPorts         []int
	allowedNetworks      []net.IPNet
	allowedMethods       []string
	disableXForwardedFor bool
}

// Option is a single option.
type Option func(*options)

// WithWhitelist sets the whitelist.
func WithWhitelist(wl []string) Option {
	return func(o *options) {
		o.whitelist = wl
	}
}

// WithBlacklist sets the blacklist.
func WithBlacklist(bl []string) Option {
	return func(o *options) {
		o.blacklist = bl
	}
}
