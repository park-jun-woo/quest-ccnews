//ff:func feature=ingestion type=helper control=sequence
//ff:what 테스트 헬퍼. 주어진 문자열을 gzip 압축한 바이트로 만든다(warc.paths.gz 응답 모킹용).

package ingest

import (
	"bytes"
	"compress/gzip"
	"testing"
)

func gzBytes(t *testing.T, s string) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write([]byte(s)); err != nil {
		t.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}
