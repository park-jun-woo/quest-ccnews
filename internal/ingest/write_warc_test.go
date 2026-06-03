//ff:func feature=ingestion type=helper control=iteration dimension=1
//ff:what 테스트 헬퍼. 주어진 (type, target-uri) 레코드들로 비압축 WARC 파일을 만들어 경로를 반환한다.

package ingest

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/slyrz/warc"
)

// writeWarc builds a small uncompressed WARC file the slyrz reader can parse.
func writeWarc(t *testing.T, records []warcRecord) string {
	t.Helper()
	var buf bytes.Buffer
	w := warc.NewWriter(&buf)
	for _, r := range records {
		rec := warc.NewRecord()
		rec.Header["warc-type"] = r.Type
		rec.Header["warc-target-uri"] = r.URI
		rec.Content = bytes.NewReader([]byte("body-bytes"))
		if _, err := w.WriteRecord(rec); err != nil {
			t.Fatalf("WriteRecord: %v", err)
		}
	}
	path := filepath.Join(t.TempDir(), "test.warc")
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}
