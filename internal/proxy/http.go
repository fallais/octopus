package proxy

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

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

func (p *Proxy) cache(u *url.URL, b io.ReadCloser) error {
	if !strings.HasSuffix(u.String(), "ico") {
		fmt.Println("do not need to cache")
		return nil
	}

	hash := sha256.Sum256([]byte(u.String()))
	body, _ := ioutil.ReadAll(b)

	return p.Cache.Set(hex.EncodeToString(hash[:]), body, time.Hour)
}
