//ff:type feature=ingestion type=model
//ff:what CC-NEWS 다운로드 클라이언트. HTTP 클라이언트, user-agent, WARC 캐시 디렉터리를 묶은 얇은 IO 핸들.

package ingest

import (
	"net/http"
)

// Client is the thin IO handle for CC-NEWS downloads: an HTTP client, the
// user-agent to send, and the directory where WARCs are cached on disk.
type Client struct {
	http      *http.Client
	userAgent string
	cacheDir  string
}
