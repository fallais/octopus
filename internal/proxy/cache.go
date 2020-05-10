package proxy

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"octopus/internal/cache"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// CacheManager is the cache manager.
type CacheManager struct {
	IsEnabled bool
	Cache     cache.Cache
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// RoundTrip is the implementation of the RoundTripper interface.
func (c *CacheManager) RoundTrip(r *http.Request) (*http.Response, error) {
	// Try to get the ressource from the cache if it is enabled
	if c.IsEnabled {
		cachedRessource, err := c.Cache.Get(generateCacheKey(r))
		if err == nil {
			logrus.Debugln("ressource is in cache")

			resp := &http.Response{
				Request:    r,
				Proto:      r.Proto,
				StatusCode: http.StatusOK,
				Header:     r.Header,
				Body:       ioutil.NopCloser(bytes.NewReader(cachedRessource)),
			}

			return resp, nil
		}
	}

	// Do the request
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		return nil, fmt.Errorf("error while doing the request: %s", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading body: %s", err)
	}

	// Set the ressource in cache if it is enabled
	if c.IsEnabled {
		go func() {
			err = c.Cache.Set(generateCacheKey(r), []byte(string(body)))
			if err != nil {
				logrus.WithError(err).Errorln("error while adding ressource to cache")
				return
			}
		}()
	}

	return resp, nil
}

func matchMime(m string) bool {
	for _, mime := range Mimes {
		if strings.HasSuffix(m, mime) {
			return true
		}
	}

	return false
}

func generateCacheKey(r *http.Request) string {
	hash := sha256.Sum256([]byte(r.URL.String()))

	return hex.EncodeToString(hash[:])
}
