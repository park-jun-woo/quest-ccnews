//ff:func feature=cli type=helper control=sequence
//ff:what 테스트 헬퍼. saveAndSweep용 전역(sessionPath/submitOut)을 임시 session.json·out.jsonl로 가리키고 종료 시 복원한 뒤 두 경로를 돌려준다.

package cmd

import (
	"path/filepath"
	"testing"
)

// setSaveAndSweepPaths points the package globals at a fresh temp session.json
// and out.jsonl, restoring both afterward, and returns those paths.
func setSaveAndSweepPaths(t *testing.T) (sessPath, outPath string) {
	t.Helper()
	dir := t.TempDir()
	sessPath = filepath.Join(dir, "session.json")
	outPath = filepath.Join(dir, "out.jsonl")
	prevSess, prevOut := sessionPath, submitOut
	sessionPath, submitOut = sessPath, outPath
	t.Cleanup(func() { sessionPath, submitOut = prevSess, prevOut })
	return sessPath, outPath
}
