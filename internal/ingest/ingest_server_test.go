//ff:func feature=ingestion type=helper control=sequence
//ff:what 테스트 헬퍼. WARC 1개를 담은 warc.paths.gz와 그 WARC를 서빙하는 테스트 서버를 만든다.

package ingest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ingestServer serves a warc.paths.gz listing one WARC, and that WARC.
func ingestServer(t *testing.T, objectPath string, warc []byte) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "warc.paths.gz"):
			w.Write(gzBytes(t, objectPath+"\n"))
		case strings.HasSuffix(r.URL.Path, ".warc.gz"):
			w.Write(warc)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}
