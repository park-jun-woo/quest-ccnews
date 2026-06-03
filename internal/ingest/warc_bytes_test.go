//ff:func feature=ingestion type=helper control=iteration dimension=1
//ff:what 테스트 헬퍼. 주어진 response URI들로 메모리 상의 비압축 WARC 바이트를 만든다.

package ingest

import (
	"bytes"
	"testing"

	"github.com/slyrz/warc"
)

// warcBytes builds an in-memory uncompressed WARC with the given response URIs.
func warcBytes(t *testing.T, uris ...string) []byte {
	t.Helper()
	var buf bytes.Buffer
	w := warc.NewWriter(&buf)
	for _, u := range uris {
		rec := warc.NewRecord()
		rec.Header["warc-type"] = "response"
		rec.Header["warc-target-uri"] = u
		rec.Content = bytes.NewReader([]byte("body"))
		if _, err := w.WriteRecord(rec); err != nil {
			t.Fatal(err)
		}
	}
	return buf.Bytes()
}
