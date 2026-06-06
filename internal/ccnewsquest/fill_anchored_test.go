//ff:func feature=gate type=helper control=iteration dimension=1
//ff:what fillAnchored 단위 테스트. present 필드의 Anchored를 checkField 결과로 채우는지 직접 검증한다. ① 유효앵커 전부 원문 substring인 필드 → true. ② 유효앵커 0개(공백)인 필드 → false. ③ 환각 앵커(원문에 없음) 필드 → false. ④ nil(부재) 필드는 건드리지 않음. ⑤ 채워진 Anchored가 같은 Source·checkField와 정확히 일치(동어반복 동치).
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestFillAnchored(t *testing.T) {
	const source = "the quick brown fox jumps over the lazy dog"

	ev := &session.Event6{
		Who:  &session.Field{Value: "x", Anchors: []string{"quick brown"}},  // substring → true
		What: &session.Field{Value: "y", Anchors: []string{"absent token"}}, // hallucination → false
		When: &session.Field{Value: "z", Anchors: []string{"   "}},          // zero valid → false
		// Where/How/Why nil — must stay nil and untouched.
	}

	fillAnchored(ev, source)

	if !ev.Who.Anchored {
		t.Errorf("who.Anchored = false, want true (anchor is a source substring)")
	}
	if ev.What.Anchored {
		t.Errorf("what.Anchored = true, want false (hallucinated anchor)")
	}
	if ev.When.Anchored {
		t.Errorf("when.Anchored = true, want false (zero valid anchors)")
	}
	if ev.Where != nil || ev.How != nil || ev.Why != nil {
		t.Errorf("absent fields mutated to non-nil: where=%v how=%v why=%v", ev.Where, ev.How, ev.Why)
	}

	// Tautological equivalence: each present field's Anchored equals checkField over
	// the SAME source (the gate's own function), so the persisted flag can never
	// contradict the verdict.
	for name, f := range map[string]*session.Field{"who": ev.Who, "what": ev.What, "when": ev.When} {
		status, _ := checkField(f, source)
		want := status == statusAnchored
		if f.Anchored != want {
			t.Errorf("%s: Anchored=%v, checkField says %v (contradiction)", name, f.Anchored, want)
		}
	}
}
