//ff:func feature=cli type=helper control=sequence
//ff:what 테스트 헬퍼. event6 JSON 문서를 임시 파일에 쓰고 그 경로를 돌려준다.

package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

// writeEvent6File writes an event6 JSON document and returns its path.
func writeEvent6File(t *testing.T, js string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "ev.json")
	if err := os.WriteFile(p, []byte(js), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}
