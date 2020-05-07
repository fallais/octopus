package headers

// HopByHopHeaders are meaningful only for a single transport-level connection, and must not be retransmitted by proxies or cached.
// Source : https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
var HopByHopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"TE",
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}
