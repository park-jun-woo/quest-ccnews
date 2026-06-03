//ff:func feature=anchor type=helper control=sequence
//ff:what 필수 필드 value가 플레이스홀더("Subject")면 앵커가 진짜여도 FAIL이고 Reason이 "플레이스홀더/공허함"과 필드명을 언급하는지 검증한다(Phase009 L3).

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_RequiredPlaceholderValueFail(t *testing.T) {
	ev := &session.Event6{
		Who:  fld("Subject", "Alice"), // placeholder value, real anchor
		When: fld("Monday", "Monday"),
		What: fld("signed treaty", "sign the treaty"),
	}
	res := Gate(ev, gateSource)
	if res.Verdict != FAIL {
		t.Fatalf("Verdict = %s, want FAIL (placeholder value)", res.Verdict)
	}
	if !strings.Contains(res.Reason, "플레이스홀더/공허함") || !strings.Contains(res.Reason, "who") {
		t.Errorf("Reason %q should name field who and the placeholder reason", res.Reason)
	}
}
