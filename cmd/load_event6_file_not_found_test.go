//ff:func feature=cli type=helper control=sequence level=error
//ff:what loadEvent6가 없는 파일 경로에서 에러를 반환하는지 검증한다.

package cmd

import (
	"path/filepath"
	"testing"
)

func TestLoadEvent6_FileNotFound(t *testing.T) {
	if _, err := loadEvent6(filepath.Join(t.TempDir(), "missing.json"), nil); err == nil {
		t.Fatal("want error for missing file")
	}
}
