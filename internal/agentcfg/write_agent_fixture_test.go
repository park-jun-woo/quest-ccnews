//ff:func feature=cli type=helper control=sequence level=error
//ff:what writeAgentFixture(t) — 캔드 WARC(HTML 응답)를 임시 캐시에 쓰고, 그 WARC를 가리키는 TODO 기사 1건을 담은 session.json을 만들어 (sessionPath, outPath, cacheDir)을 돌려준다. agent 통합 테스트의 무네트워크 픽스처(공유 상수 agentPassHTML 사용).
package agentcfg

import (
	"bytes"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
	"github.com/slyrz/warc"
)

// writeAgentFixture writes a candy WARC + a session.json holding one TODO article
// pointing at it, and returns the session path, out path, and cache dir.
func writeAgentFixture(t *testing.T) (sessionPath, outPath, cacheDir string) {
	t.Helper()
	cacheDir = t.TempDir()

	body := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		"Content-Length: " + strconv.Itoa(len(agentPassHTML)) + "\r\n" +
		"\r\n" +
		agentPassHTML
	var buf bytes.Buffer
	w := warc.NewWriter(&buf)
	rec := warc.NewRecord()
	rec.Header["warc-type"] = "response"
	rec.Header["warc-target-uri"] = "https://example.com/a"
	rec.Content = bytes.NewReader([]byte(body))
	if _, err := w.WriteRecord(rec); err != nil {
		t.Fatalf("WriteRecord: %v", err)
	}
	file := "CC-NEWS-test.warc"
	if err := os.WriteFile(filepath.Join(cacheDir, file), buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}

	s := quest.New()
	it := &quest.Item{Key: "https://example.com/a", State: quest.TODO}
	if err := it.SetPayload(&session.Article{
		URL:   "https://example.com/a",
		Host:  "example.com",
		State: session.TODO,
		WARC:  &session.WARCLoc{File: file, Offset: 0},
	}); err != nil {
		t.Fatal(err)
	}
	s.Items = []*quest.Item{it}

	dir := t.TempDir()
	sessionPath = filepath.Join(dir, "session.json")
	outPath = filepath.Join(dir, "out.jsonl")
	if err := s.Save(sessionPath); err != nil {
		t.Fatalf("Save session: %v", err)
	}
	return sessionPath, outPath, cacheDir
}
