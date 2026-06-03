//ff:func feature=anchor type=helper control=sequence
//ff:what MaxTries-1회 실패 후 한 번 더 FAIL이면 State=DONE으로 잠기고 Verdict/Reason 기록·CollectedAt 비움·마지막 Try=MaxTries인지 검증한다.

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestApply_FailExhaustToDone(t *testing.T) {
	// Already failed MaxTries-1 times; one more FAIL locks to DONE.
	a := &session.Article{State: session.TODO, Tries: session.MaxTries - 1}
	res := Result{Verdict: FAIL, Reason: "final bad"}
	Apply(a, &session.Event6{}, res, "ts")

	if a.State != session.DONE {
		t.Errorf("State = %s, want DONE", a.State)
	}
	if a.Tries != session.MaxTries {
		t.Errorf("Tries = %d, want %d", a.Tries, session.MaxTries)
	}
	if a.Verdict != "FAIL" || a.VerdictReason != "final bad" {
		t.Errorf("Verdict=%q Reason=%q", a.Verdict, a.VerdictReason)
	}
	if a.CollectedAt != "" {
		t.Errorf("CollectedAt = %q, want empty on FAIL→DONE", a.CollectedAt)
	}
	if a.Log[len(a.Log)-1].Try != session.MaxTries {
		t.Errorf("last Try = %d, want %d", a.Log[len(a.Log)-1].Try, session.MaxTries)
	}
}
