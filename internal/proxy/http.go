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
	p.updateXFFHeader(r.Header)

	// Clean the RequestURI
	r.RequestURI = ""

	// Do the request
	resp, err := p.httpClient.Do(r)
	if err != nil {
		logrus.WithError(err).Errorln("error while doing the request")
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// Copy headers
	p.copyHeader(w.Header(), resp.Header)

	// Remove hop-by-hop headers
	p.removeHopByHopHeaders(r.Header)

	// Write the status code
	w.WriteHeader(resp.StatusCode)

	// Write the response
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		logrus.WithError(err).Errorln("error while copying the stream")
		return
	}

}
