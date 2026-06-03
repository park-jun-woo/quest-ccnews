//ff:func feature=robots type=helper control=sequence
//ff:what 매칭되는 그룹이 없으면 Evaluate가 default-allow를 반환하는지 검증한다.

package robots

import "testing"

func TestEvaluateNoGroupDefaultAllow(t *testing.T) {
	rs := &Ruleset{Groups: []Group{{Agents: []string{"otherbot"}}}}
	d := Evaluate(rs, "parkjunwoo-quest", "/x")
	if !d.Allowed || d.Rule != "" {
		t.Errorf("no matching group → default-allow, got %+v", d)
	}
}
