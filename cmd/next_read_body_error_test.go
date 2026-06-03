//ff:func feature=cli type=helper control=sequence level=error
//ff:what runNext가 기사가 가리키는 WARC 파일이 캐시에 없어 원문 재독이 실패할 때 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunNext_ReadBodyError(t *testing.T) {
	// Article references a WARC file that does not exist in the cache dir.
	p := writeSessionWith(t, "nonexistent.warc")
	sessionPath = p
	prevCache := nextCacheDir
	nextCacheDir = t.TempDir()
	t.Cleanup(func() { sessionPath = "session.json"; nextCacheDir = prevCache })

	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runNext(cmd, nil); err == nil {
		t.Fatal("want error from missing WARC body")
	}
}
