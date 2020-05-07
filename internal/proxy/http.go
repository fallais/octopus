package proxy

import (
	"fmt"
	"io"
	"net/http"
)

// HTTPHandler for HTTP connections.
func (p *Proxy) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	// Remove hop-by-hop headers
	p.removeHopByHopHeaders(r.Header)

	// Append the XFF to the other XFF
	p.updateXFFHeader(r.Header, r.Host)

	fmt.Println(r.Header)

	// TODO : processHeaders

	// Do the request
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// Copy headers
	p.copyHeader(w.Header(), resp.Header)

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}
