package proxy

import (
	"io"
	"net"
	"net/http"
	"time"

	"octopus/internal/metrics"

	"github.com/sirupsen/logrus"
)

// Handler is the HTTP handler.
func (p *Proxy) Handler(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"remote_addr": r.RemoteAddr,
		"method":      r.Method,
		"url":         r.URL,
	}).Debugln("Request incoming")

	go metrics.IncrementRequestCount(r.Method)

	// Blacklist
	for _, site := range p.opts.blacklist {
		if site == r.URL.Host {
			logrus.WithFields(logrus.Fields{
				"remote_addr": r.RemoteAddr,
				"method":      r.Method,
				"url":         r.URL,
			}).Debugln("URL is blacklisted")

			p.blacklist(w, r)

			// TODO : go metrics.IncrementBlacklistedCount(r.Method)

			return
		}
	}

	if r.Method == http.MethodConnect {
		p.https(w, r)
	} else {
		p.http(w, r)
	}
}

func (p *Proxy) http(w http.ResponseWriter, r *http.Request) {
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

func (p *Proxy) https(w http.ResponseWriter, r *http.Request) {
	destConn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}

func (p *Proxy) blacklist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL is blacklisted"))
}
