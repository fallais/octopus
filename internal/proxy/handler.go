package proxy

import (
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
	for _, site := range p.opts.blacklist {
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

	if r.Method == http.MethodConnect {
		p.HTTPSHandler(w, r)
	} else {
		p.HTTPHandler(w, r)
	}
}
