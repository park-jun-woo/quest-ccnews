//ff:func feature=anchor type=helper control=sequence
//ff:what present 선택 필드 value가 짧음("a")이면 앵커가 진짜여도 value 위생으로 FAIL이고 Reason이 "플레이스홀더/공허함"과 필드명을 언급하는지 검증한다(Phase009 L3).

package anchor

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestGate_OptionalPlaceholderValueFail(t *testing.T) {
	ev := &session.Event6{
		Who:   fld("Alice", "Alice"),
		When:  fld("Monday", "Monday"),
		What:  fld("signed treaty", "sign the treaty"),
		Where: fld("a", "Paris"), // 1-rune value, real anchor → value hygiene FAIL
	}
	res := Gate(ev, gateSource)
	if res.Verdict != FAIL {
		t.Fatalf("Verdict = %s, want FAIL (too-short optional value)", res.Verdict)
	}
	if !strings.Contains(res.Reason, "플레이스홀더/공허함") || !strings.Contains(res.Reason, "where") {
		t.Errorf("Reason %q should name field where and the hygiene reason", res.Reason)
	}
}
