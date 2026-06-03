//ff:func feature=output type=helper control=sequence level=error
//ff:what 파일의 비어있지 않은 줄 수를 세는 테스트 헬퍼. 파일이 없으면 0, 내용이 비면 0. 후행 개행은 무시한다.

package output

import (
	"os"
	"strings"
	"testing"
)

func countLines(t *testing.T, path string) int {
	t.Helper()
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return 0
	}
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	trimmed := strings.TrimRight(string(data), "\n")
	if trimmed == "" {
		return 0
	}
	return len(strings.Split(trimmed, "\n"))
}
