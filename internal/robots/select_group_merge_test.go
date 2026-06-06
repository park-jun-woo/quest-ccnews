//ff:func feature=robots type=helper control=sequence
//ff:what 동일 * 에이전트 레코드 둘(빈 Disallow + Disallow: /reporter/)을 병합해 RFC 9309 §2.2.1대로 평가되는지 검증한다.

package robots

import "testing"

func TestSelectGroupMergeSameAgent(t *testing.T) {
	rs := Parse([]byte(
		"User-agent: *\n" +
			"Disallow:\n" +
			"\n" +
			"User-agent: *\n" +
			"Disallow: /reporter/\n",
	))

	if d := Evaluate(rs, "parkjunwoo-quest", "/reporter/x"); d.Allowed {
		t.Errorf("/reporter/x should be blocked after merge, got %+v", d)
	}
	if d := Evaluate(rs, "parkjunwoo-quest", "/6993730"); !d.Allowed {
		t.Errorf("/6993730 should be allowed after merge, got %+v", d)
	}
}
