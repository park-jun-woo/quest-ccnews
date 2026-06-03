//ff:func feature=ingestion type=helper control=sequence
//ff:what 기본 타임아웃을 가진 CC-NEWS 다운로드 클라이언트를 만든다. cacheDir는 받은 .warc.gz 캐시 위치.

package ingest

import (
	"net/http"
	"time"
)

// NewClient builds a download client with a sane default timeout. cacheDir is
// where downloaded .warc.gz files are stored (retained while TODO articles for
// that WARC remain — Phase003 §열린결정).
func NewClient(userAgent, cacheDir string) *Client {
	return &Client{
		http:      &http.Client{Timeout: 30 * time.Minute},
		userAgent: userAgent,
		cacheDir:  cacheDir,
	}
}
