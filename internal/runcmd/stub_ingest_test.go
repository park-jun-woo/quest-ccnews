//ff:func feature=ingestion type=command control=sequence
//ff:what stubIngest 테스트 헬퍼. ingestRun 패키지 변수를 무네트워크 fn으로 테스트 수명 동안 교체하고 Cleanup으로 원복해, run 테스트가 실제 다운로드 없이 ingest 루프 동작을 흉내내게 한다.

package runcmd

import (
	"io"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/ingest"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// stubIngest swaps the ingestRun package var (no-network) for the test's lifetime.
func stubIngest(t *testing.T, fn func(*session.Session, ingest.RunOptions) error) {
	t.Helper()
	orig := ingestRun
	ingestRun = func(c *ingest.Client, s *session.Session, opt ingest.RunOptions, w io.Writer) error {
		return fn(s, opt)
	}
	t.Cleanup(func() { ingestRun = orig })
}
