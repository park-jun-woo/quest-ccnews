//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what 빈 Disallow만 있는 룰셋(allow all 관용구)이 임의 path를 차단하지 않고 허용하는지 검증한다(D1 회귀 방지).

package robots

import "testing"

func TestEvaluateEmptyDisallowAllowsAll(t *testing.T) {
	rs := Parse([]byte("User-agent: *\nDisallow:\n"))
	for _, path := range []string{"/", "/6993730", "/anything/deep/here"} {
		d := Evaluate(rs, "parkjunwoo-quest", path)
		if !d.Allowed {
			t.Errorf("path %q: expected Allowed, got %+v", path, d)
		}
	}
}
