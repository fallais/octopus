package proxy

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"remote_addr": r.RemoteAddr,
		"method":      r.Method,
		"url":         r.URL,
	}).Debugln("Request incoming")

	// Blacklist
	for _, site := range p.Blacklist {
		if site == r.URL.Host {
			logrus.WithFields(logrus.Fields{
				"remote_addr": r.RemoteAddr,
				"method":      r.Method,
				"url":         r.URL,
			}).Debugln("URL is blacklisted")

			p.BlacklistedHandler(w, r)

			return
		}
	}

	// Check the cache
	if p.IsCacheEnabled {
		// Generate the hash of the ressource
		hash := sha256.Sum256([]byte(r.URL.String()))

		// Get the object from the cache
		object, err := p.Cache.Get(hex.EncodeToString(hash[:]))
		if err == nil {
			w.WriteHeader(http.StatusResetContent)

			io.Copy(w, bytes.NewReader(object))

			return
		}

		logrus.Debugln("Object is not in cache")
	}

	if r.Method == http.MethodConnect {
		p.HTTPSHandler(w, r)
	} else {
		p.HTTPHandler(w, r)
	}
}
