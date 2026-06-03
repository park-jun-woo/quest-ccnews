//ff:func feature=robots type=helper control=sequence
//ff:what 이미 BLOCKED인 기사에 대해 Gate가 false를 반환하고 상태를 바꾸지 않는지 검증한다.

package robots

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGateLockedBlockedReturnsFalse(t *testing.T) {
	a := &session.Article{State: session.BLOCKED, URL: "https://example.com/public"}
	if Gate(a, disallowRuleset(), "parkjunwoo-quest") {
		t.Errorf("already BLOCKED article should return false")
	}
	if a.State != session.BLOCKED {
		t.Errorf("state must not change, got %v", a.State)
	}
}
