//ff:func feature=gate type=helper control=sequence
//ff:what Def().Rules() 카탈로그 크기 스모크. Def가 앵커 규칙 6개를 내는지 검증한다.

package ccnewsquest

import "testing"

func TestDefRulesCount(t *testing.T) {
	if n := len(Def("ua", "cache").Rules()); n != 6 {
		t.Fatalf("Rules count = %d, want 6", n)
	}
}
