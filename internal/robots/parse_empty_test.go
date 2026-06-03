//ff:func feature=robots type=helper control=sequence
//ff:what 빈 내용에 대해 Parse가 그룹 없는 룰셋을 반환하는지 검증한다.

package robots

import "testing"

func TestParseEmptyContent(t *testing.T) {
	rs := Parse([]byte(""))
	if len(rs.Groups) != 0 {
		t.Errorf("empty content should yield no groups, got %d", len(rs.Groups))
	}
}
