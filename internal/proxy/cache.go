package proxy

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)

func (p *Proxy) cache(u *url.URL, b io.ReadCloser) error {
	if !matchMime(u.String()) {
		fmt.Println("do not need to cache")
		return nil
	}

	hash := sha256.Sum256([]byte(u.String()))
	body, _ := ioutil.ReadAll(b)

	return p.Cache.Set(hex.EncodeToString(hash[:]), body, time.Hour)
}

func matchMime(m string) bool {
	for _, mime := range Mimes {
		if strings.HasSuffix(m, mime) {
			return true
		}
	}

	return false
}
