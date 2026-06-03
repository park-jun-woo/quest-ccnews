//ff:func feature=cli type=helper control=sequence
//ff:what 테스트 헬퍼. runIngestion 테스트용으로 out/cache 플래그를 임시 디렉터리로 잡고, 종료 시 전역 플래그와 ingestRun seam을 원복한다.

package cmd

import (
	"path/filepath"
	"testing"
)

// withRunStub points the run command's out/cache flags into dir and restores the
// run globals and the ingestRun seam on cleanup, so each runIngestion test runs
// in isolation without touching the real CC-NEWS network.
func withRunStub(t *testing.T, dir string) {
	t.Helper()
	prevTrack, prevOut, prevCache := runTrack, runOut, runCacheDir
	prevRun := ingestRun
	runOut = filepath.Join(dir, "out.jsonl")
	runCacheDir = dir
	t.Cleanup(func() {
		runTrack, runOut, runCacheDir = prevTrack, prevOut, prevCache
		ingestRun = prevRun
		sessionPath = "session.json"
	})
}
