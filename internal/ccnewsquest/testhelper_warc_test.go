//ff:func feature=gate type=helper control=sequence
//ff:what Prepare 테스트 헬퍼. 한 HTML 본문을 WARC response 레코드(전체 HTTP 응답)로 감싼 비압축 .warc를 임시 캐시에 쓰고 (cacheDir, file)을 돌려준다. ingest 패키지의 동명 헬퍼와 같은 형식(네트워크·실 WARC 불요).

package ccnewsquest

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/slyrz/warc"
)

// writeWarcHTML writes an uncompressed WARC whose single response record carries
// the given HTML wrapped in a full HTTP/1.1 response, mirroring how CC-NEWS wraps
// origin responses. It returns the cache dir and file name for ingest ReadBody.
func writeWarcHTML(t *testing.T, html string) (cacheDir, file string) {
	t.Helper()
	body := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Length: " + strconv.Itoa(len(html)) + "\r\n" +
		"\r\n" +
		html

	var buf bytes.Buffer
	w := warc.NewWriter(&buf)
	rec := warc.NewRecord()
	rec.Header["warc-type"] = "response"
	rec.Header["warc-target-uri"] = "https://example.com/a"
	rec.Content = bytes.NewReader([]byte(body))
	if _, err := w.WriteRecord(rec); err != nil {
		t.Fatalf("WriteRecord: %v", err)
	}

	cacheDir = t.TempDir()
	file = "CC-NEWS-test.warc"
	if err := os.WriteFile(filepath.Join(cacheDir, file), buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	return cacheDir, file
}

// passHTML is an article page that clears the trust gate (structured JSON-LD
// source + self-declared title + body >= MinBodyLen), so extract.Apply PASSes.
const passHTML = `<html><head>` +
	`<script type="application/ld+json">{"@type":"NewsArticle","headline":"Framework Agreement","author":{"name":"Reporter"},"datePublished":"2026-06-05","articleBody":"` +
	passBody +
	`","inLanguage":"en"}</script>` +
	`</head><body></body></html>`

// passBody is comfortably above MinBodyLen (200 bytes).
const passBody = "This is a sufficiently long article body that exceeds the minimum length " +
	"threshold required by the trust gate so that the article passes the body-length " +
	"check and is not skipped as too short for anchoring purposes. It runs on for a while."

// skipHTML has no structured data, so the trust gate FAILs (SkipNoStructured).
const skipHTML = `<html><body>no structured data here at all</body></html>`
