//ff:func feature=output type=helper control=selection dimension=1
//ff:what 종단 기사+호스트를 JSONL 레코드로 렌더링하는 순수 함수. PASS/REVIEW는 renderCollected, BLOCKED/SKIPPED/DONE은 renderAudit, 종단이 아니면 nil.

package output

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// Render builds the JSONL Record for a terminal-state article. host may be nil
// (the host cache entry is optional). It is pure: no IO, no mutation. A non-
// terminal article (or unknown state) returns nil so callers skip it.
func Render(a *session.Article, host *session.Host) *Record {
	if a == nil {
		return nil
	}
	switch a.State {
	case session.PASS, session.REVIEW:
		return renderCollected(a, host)
	case session.BLOCKED, session.SKIPPED, session.DONE:
		return renderAudit(a, host)
	default:
		return nil
	}
}
