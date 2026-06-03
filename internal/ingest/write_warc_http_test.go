//ff:func feature=ingestion type=helper control=iteration dimension=1
//ff:what 테스트 헬퍼. response 레코드마다 전체 HTTP 응답을 담은 비압축 WARC를 캐시 디렉터리에 쓰고 (cacheDir, file)을 돌려준다.

package ingest

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/slyrz/warc"
)

// writeWarcHTTP writes an uncompressed WARC whose response records carry full
// HTTP responses, returning the cache dir and file name for ReadBody.
func writeWarcHTTP(t *testing.T, bodies ...string) (cacheDir, file string) {
	t.Helper()
	var buf bytes.Buffer
	w := warc.NewWriter(&buf)
	for i, body := range bodies {
		rec := warc.NewRecord()
		rec.Header["warc-type"] = "response"
		rec.Header["warc-target-uri"] = "https://example.com/" + itoa(i)
		rec.Content = bytes.NewReader([]byte(httpResponse(body)))
		if _, err := w.WriteRecord(rec); err != nil {
			t.Fatalf("WriteRecord: %v", err)
		}
	}
	cacheDir = t.TempDir()
	file = "CC-NEWS-test.warc"
	if err := os.WriteFile(filepath.Join(cacheDir, file), buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	return cacheDir, file
}
