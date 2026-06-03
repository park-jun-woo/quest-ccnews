//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what BLOCKED이 아닌 잠긴 상태(PASS/REVIEW/DONE/SKIPPED)에 대해 Gate가 true를 반환하고 상태를 보존하는지 검증한다.

package robots

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGateLockedNonBlockedReturnsTrue(t *testing.T) {
	for _, st := range []session.State{session.PASS, session.REVIEW, session.DONE, session.SKIPPED} {
		a := &session.Article{State: st, URL: "https://example.com/secret/x"}
		if !Gate(a, disallowRuleset(), "parkjunwoo-quest") {
			t.Errorf("locked non-BLOCKED state %v should return true", st)
		}
		if a.State != st {
			t.Errorf("state %v must not change", st)
		}
	}
}
