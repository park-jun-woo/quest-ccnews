//ff:func feature=cli type=helper control=sequence
//ff:what 테스트 헬퍼. HTTP 응답 1건을 담은 WARC를 새 캐시 디렉터리에 쓰고 (cacheDir, file)을 돌려준다(레코드는 offset 0).

package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/slyrz/warc"
)

// writeCacheWarc writes a one-record WARC (full HTTP response) into a fresh
// cache dir and returns (cacheDir, file). The record is at ordinal/offset 0.
func writeCacheWarc(t *testing.T, html string) (cacheDir, file string) {
	t.Helper()
	var buf bytes.Buffer
	w := warc.NewWriter(&buf)
	rec := warc.NewRecord()
	rec.Header["warc-type"] = "response"
	rec.Header["warc-target-uri"] = "https://example.com/a"
	rec.Content = bytes.NewReader([]byte(httpResp(html)))
	if _, err := w.WriteRecord(rec); err != nil {
		t.Fatal(err)
	}
	cacheDir = t.TempDir()
	file = "CC-NEWS-test.warc"
	if err := os.WriteFile(filepath.Join(cacheDir, file), buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	return cacheDir, file
}
