package proxy

import (
	"io"
	"net/http"
)

// HTTPHandler for HTTP connections.
func (proxy *Proxy) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	delHopHeaders(r.Header)
	copyHeader(w.Header(), resp.Header)

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}
