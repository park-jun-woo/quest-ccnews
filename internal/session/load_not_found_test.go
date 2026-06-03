//ff:func feature=session type=helper control=sequence
//ff:what Load가 존재하지 않는 파일에 대해 에러를 반환하는지 검증한다.

package session

import (
	"path/filepath"
	"testing"
)

func TestLoadNotFound(t *testing.T) {
	if _, err := Load(filepath.Join(t.TempDir(), "nope.json")); err == nil {
		t.Error("Load of missing file expected error, got nil")
	}
}
