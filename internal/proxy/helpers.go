package proxy

import (
	"io"
	"net/http"
	"strings"

	"octopus/internal/headers"
)

func (p *Proxy) copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			if v == "X-Forwarded-For" && p.DisableForwardedFor {
				continue
			}

			dst.Add(k, v)
		}
	}
}

func (p *Proxy) removeHopByHopHeaders(header http.Header) {
	for _, h := range headers.HopByHopHeaders {
		header.Del(h)
	}
}

func (p *Proxy) updateXFFHeader(header http.Header, host string) {
	if prior, ok := header["X-Forwarded-For"]; ok {
		host = strings.Join(prior, ", ") + ", " + host
	}

	header.Set("X-Forwarded-For", host)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}
