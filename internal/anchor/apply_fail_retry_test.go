//ff:func feature=anchor type=helper control=sequence
//ff:what Apply(FAIL, 첫 시도)가 State=TODO 유지·Tries++·CollectedAt/Verdict 비움·Log 1건(FAIL,try=1)을 남기는지 검증한다.

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestApply_FailRetry(t *testing.T) {
	a := &session.Article{State: session.TODO}
	res := Result{Verdict: FAIL, Reason: "bad"}
	Apply(a, &session.Event6{}, res, "ts")

	if a.State != session.TODO {
		t.Errorf("State = %s, want TODO (retry)", a.State)
	}
	if a.Tries != 1 {
		t.Errorf("Tries = %d, want 1", a.Tries)
	}
	if a.CollectedAt != "" {
		t.Errorf("CollectedAt = %q, want empty on FAIL", a.CollectedAt)
	}
	if a.Verdict != "" {
		t.Errorf("Verdict = %q, want empty while retrying", a.Verdict)
	}
	if len(a.Log) != 1 || a.Log[0].Verdict != "FAIL" || a.Log[0].Try != 1 {
		t.Errorf("Log = %+v, want one FAIL attempt at try 1", a.Log)
	}
}
