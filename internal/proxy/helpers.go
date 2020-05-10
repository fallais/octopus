package proxy

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"

	"octopus/internal/proxy/headers"
)

func (p *Proxy) copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			if v == "X-Forwarded-For" && p.opts.disableXForwardedFor {
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

func (p *Proxy) updateXFFHeader(header http.Header) {
	hn, err := os.Hostname()
	if err != nil {
		hn = "Octopus"
	}

	xff, ok := header["X-Forwarded-For"]
	if ok {
		hn = strings.Join(xff, ", ") + ", " + hn
	}

	header.Set("X-Forwarded-For", hn)
}

// TODO : Via Header !
func (p *Proxy) updateViaHeader(header http.Header) {
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func isGziped(r bufio.Reader) (bool, error) {
	firstTwoBytes, err := r.Peek(2)
	if err != nil {
		return false, err
	}

	return firstTwoBytes[0] == 31 && firstTwoBytes[1] == 139, nil
}
