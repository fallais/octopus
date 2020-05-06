package proxy

import (
	"net/http"
)

// BlacklistedHandler for blacklisted URL.
func (proxy *Proxy) BlacklistedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL is blacklisted"))
}
