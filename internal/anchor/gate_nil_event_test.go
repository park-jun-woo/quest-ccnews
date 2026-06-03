//ff:func feature=anchor type=helper control=sequence
//ff:what Gate(nil)이 FAIL과 비어있지 않은 Reason을 반환하는지 검증한다.

package anchor

import "testing"

func TestGate_NilEvent(t *testing.T) {
	res := Gate(nil, gateSource)
	if res.Verdict != FAIL {
		t.Fatalf("Verdict = %s, want FAIL", res.Verdict)
	}
	if res.Reason == "" {
		t.Error("Reason should be non-empty")
	}
}
