//ff:func feature=anchor type=helper control=sequence
//ff:what 연속 FAIL 두 번에서 Log가 2건이고 Try 번호가 1,2로 증가하는지 검증한다.

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestApply_LogTryCountsAcrossAttempts(t *testing.T) {
	a := &session.Article{State: session.TODO}
	Apply(a, &session.Event6{}, Result{Verdict: FAIL, Reason: "1"}, "ts")
	Apply(a, &session.Event6{}, Result{Verdict: FAIL, Reason: "2"}, "ts")
	if len(a.Log) != 2 {
		t.Fatalf("Log len = %d, want 2", len(a.Log))
	}
	if a.Log[0].Try != 1 || a.Log[1].Try != 2 {
		t.Errorf("Try numbers = %d,%d want 1,2", a.Log[0].Try, a.Log[1].Try)
	}
}
