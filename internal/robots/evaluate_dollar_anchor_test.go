//ff:func feature=robots type=helper control=sequence
//ff:what 끝 $ 앵커 패턴(/x$)이 정확히 /x만 막고 /xy는 막지 않는지 Evaluate로 검증한다.

package robots

import "testing"

func TestEvaluateDollarAnchor(t *testing.T) {
	rs := &Ruleset{Groups: []Group{{Agents: []string{"*"}, Rules: []Rule{
		{Allow: false, Pattern: "/x$"},
	}}}}
	if d := Evaluate(rs, "parkjunwoo-quest", "/x"); d.Allowed {
		t.Errorf("/x$ should block /x, got %+v", d)
	}
	if d := Evaluate(rs, "parkjunwoo-quest", "/xy"); !d.Allowed {
		t.Errorf("/x$ should not block /xy, got %+v", d)
	}
}
