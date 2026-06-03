//ff:func feature=robots type=helper control=sequence
//ff:what Gate가 Disallow 경로의 TODO 기사를 BLOCKED로 잠그고 skip_reason을 기록하는지 검증한다.

package robots

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGateBlocksTODO(t *testing.T) {
	a := &session.Article{State: session.TODO, URL: "https://example.com/secret/x"}
	allowed := Gate(a, disallowRuleset(), "parkjunwoo-quest")
	if allowed {
		t.Errorf("disallowed path should not pass")
	}
	if a.State != session.BLOCKED {
		t.Errorf("blocked article state = %v, want BLOCKED", a.State)
	}
	if a.SkipReason != "robots Disallow: /secret" {
		t.Errorf("skip reason = %q, want %q", a.SkipReason, "robots Disallow: /secret")
	}
}
