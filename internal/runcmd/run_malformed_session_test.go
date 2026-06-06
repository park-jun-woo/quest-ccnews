//ff:func feature=ingestion type=command control=sequence level=error
//ff:what run 단위테스트(손상된 세션). 잘못된 JSON 세션 파일이면 quest.Load 에러가 run에서 전파된다.

package runcmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunMalformedSessionLoadError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "session.json")
	if err := os.WriteFile(path, []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	o, cmd := newOptions(t, path)

	if err := o.run(cmd, nil); err == nil {
		t.Errorf("run() err = nil, want unmarshal error")
	}
}
