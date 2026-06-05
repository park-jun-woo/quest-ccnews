//ff:func feature=cli type=helper control=sequence level=error
//ff:what 세션을 저장하고 종단 미emit 기사를 out으로 sweep한 뒤, emit이 있었으면 Emitted 플래그를 재저장한다. submit의 SKIPPED 단락과 PASS/REVIEW/FAIL 경로가 동일하게 영속화하도록 공유.

package cmd

import (
	"github.com/park-jun-woo/quest-ccnews/internal/output"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// saveAndSweep persists the session, sweeps terminal not-yet-emitted articles to
// the JSONL output (uniform across all terminal states), then re-saves the
// Emitted flags when anything was emitted. Shared by the SKIPPED short-circuit
// and the PASS/REVIEW/FAIL path so both persist identically.
func saveAndSweep(s *session.Session) error {
	if err := s.Save(sessionPath); err != nil {
		return err
	}
	n, err := output.Sweep(s, submitOut)
	if err != nil {
		return err
	}
	if n > 0 {
		return s.Save(sessionPath)
	}
	return nil
}
