//ff:func feature=robots type=helper control=sequence
//ff:what Gate가 허용 경로의 TODO 기사를 통과시키고 상태·skip_reason을 건드리지 않는지 검증한다.

package robots

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGateAllowsTODO(t *testing.T) {
	a := &session.Article{State: session.TODO, URL: "https://example.com/public"}
	allowed := Gate(a, disallowRuleset(), "parkjunwoo-quest")
	if !allowed {
		t.Errorf("allowed path should pass, got allowed=%v", allowed)
	}
	if a.State != session.TODO {
		t.Errorf("allowed article state should stay TODO, got %v", a.State)
	}
	if a.SkipReason != "" {
		t.Errorf("allowed article should have no skip reason, got %q", a.SkipReason)
	}
}
