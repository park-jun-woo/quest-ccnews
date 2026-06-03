//ff:func feature=cli type=helper control=sequence
//ff:what runNext가 캐시 WARC에서 원문을 재독·추출해 URL과 본문 텍스트를 출력하는 정상 경로를 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunNext_HappyPath(t *testing.T) {
	cache, file := writeCacheWarc(t, "<html><body>Alice met Bob in Paris</body></html>")
	p := writeSessionWith(t, file)
	sessionPath = p
	prevCache := nextCacheDir
	nextCacheDir = cache
	t.Cleanup(func() { sessionPath = "session.json"; nextCacheDir = prevCache })

	cmd := &cobra.Command{}
	var out bytes.Buffer
	cmd.SetOut(&out)
	if err := runNext(cmd, nil); err != nil {
		t.Fatalf("runNext: %v", err)
	}
	o := out.String()
	if !strings.Contains(o, "https://example.com/a") || !strings.Contains(o, "Alice met Bob in Paris") {
		t.Errorf("output missing expected content:\n%s", o)
	}
}
