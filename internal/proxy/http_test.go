package proxy_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"octopus/internal/proxy"
)

func BenchmarkHTTPHandler(b *testing.B) {
	// Create the proxy
	p, err := proxy.NewProxy(nil, nil)
	if err != nil {
		b.Fatalf("Error while creating the proxy: %s", err)
	}

	req, err := http.NewRequest("GET", "www.test.com", nil)
	if err != nil {
		b.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		p.ServeHTTP(rr, req)
	}
}
