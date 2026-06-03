//ff:func feature=anchor type=helper control=sequence
//ff:what Apply(PASS)가 State=PASS·Verdict·Reason·CollectedAt·Event6 부착, Tries 불변, Log 1건(try=1,PASS)을 기록하는지 검증한다.

package anchor

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestApply_Pass(t *testing.T) {
	a := &session.Article{State: session.TODO}
	ev := &session.Event6{}
	res := Result{Verdict: PASS, Reason: "ok"}
	Apply(a, ev, res, "2026-06-03T00:00:00Z")

	if a.State != session.PASS {
		t.Errorf("State = %s, want PASS", a.State)
	}
	if a.Verdict != "PASS" {
		t.Errorf("Verdict = %q, want PASS", a.Verdict)
	}
	if a.VerdictReason != "ok" {
		t.Errorf("VerdictReason = %q, want ok", a.VerdictReason)
	}
	if a.CollectedAt != "2026-06-03T00:00:00Z" {
		t.Errorf("CollectedAt = %q", a.CollectedAt)
	}
	if a.Event6 != ev {
		t.Error("Event6 not attached")
	}
	if a.Tries != 0 {
		t.Errorf("Tries = %d, want 0", a.Tries)
	}
	if len(a.Log) != 1 || a.Log[0].Try != 1 || a.Log[0].Verdict != "PASS" || a.Log[0].Reason != "ok" {
		t.Errorf("Log = %+v, want one PASS attempt", a.Log)
	}
}
