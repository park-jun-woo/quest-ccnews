//ff:func feature=gate type=helper control=sequence
//ff:what Prepare 검증 공용 헬퍼. 한 필드의 payload anchored가 같은 Source 위에서 게이트 동일 함수 checkField가 내리는 statusAnchored 판정과 정확히 일치하는지 단언한다(verdict 모순 0). prepare_anchored_test의 ③ 검증 단계를 깊이 1 루프 없이 필드별로 호출하기 위한 추출.
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// assertAnchoredMatchesCheckField asserts that f.Anchored equals the gate's own
// checkField verdict over the same source — i.e. the payload anchored flag and the
// gate agree with zero contradictions.
func assertAnchoredMatchesCheckField(t *testing.T, name string, f *session.Field, source string) {
	t.Helper()
	status, _ := checkField(f, source)
	want := status == statusAnchored
	if f.Anchored != want {
		t.Errorf("%s: payload anchored=%v, checkField says %v (verdict contradiction)", name, f.Anchored, want)
	}
}
