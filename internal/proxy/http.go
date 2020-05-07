package proxy

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// HTTPHandler for HTTP connections.
func (p *Proxy) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	// Remove hop-by-hop headers
	p.removeHopByHopHeaders(r.Header)

	// Append the XFF to the other XFF
	p.updateXFFHeader(r.Header, r.Host)

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

	// Generate the hash of the ressource
	err = p.cache(r.URL, resp.Body)
	if err != nil {
		logrus.WithError(err).Errorln("Error while writing in cache")
	}
}
