//ff:func feature=anchor type=helper control=sequence
//ff:what Apply(REVIEW)가 State=REVIEW·Verdict·Reason·CollectedAt를 기록하고 Tries는 불변인지 검증한다.

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestApply_Review(t *testing.T) {
	a := &session.Article{State: session.TODO}
	res := Result{Verdict: REVIEW, Reason: "needs human"}
	Apply(a, &session.Event6{}, res, "ts")

	if a.State != session.REVIEW {
		t.Errorf("State = %s, want REVIEW", a.State)
	}
	if a.Verdict != "REVIEW" || a.VerdictReason != "needs human" {
		t.Errorf("Verdict=%q Reason=%q", a.Verdict, a.VerdictReason)
	}
	if a.CollectedAt != "ts" {
		t.Errorf("CollectedAt = %q, want ts", a.CollectedAt)
	}
	if a.Tries != 0 {
		t.Errorf("Tries = %d, want 0", a.Tries)
	}
}
