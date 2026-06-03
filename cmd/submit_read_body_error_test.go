//ff:func feature=cli type=helper control=sequence level=error
//ff:what runSubmit가 기사 WARC가 캐시에 없어 원문 재독이 실패할 때 "원문 재독 실패" 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunSubmit_ReadBodyError(t *testing.T) {
	resetSubmitFlags(t)
	p := writeSessionWith(t, "nonexistent.warc")
	evPath := writeEvent6File(t, `{"who":{"value":"A","anchors":["A"]}}`)
	submitURL = "https://example.com/a"
	submitEvent6 = evPath
	sessionPath = p
	prev := submitCacheDir
	submitCacheDir = t.TempDir()
	t.Cleanup(func() { submitCacheDir = prev })
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runSubmit(cmd, nil); err == nil || !strings.Contains(err.Error(), "원문 재독 실패") {
		t.Fatalf("want ReadBody error, got %v", err)
	}
}
