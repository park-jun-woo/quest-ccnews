//ff:func feature=robots type=helper control=sequence
//ff:what 빈 값 Disallow/Allow 지시자가 빈-패턴 규칙으로 추가되지 않는지(RFC 9309: 빈 값=효력 없음) 검증한다. D1 대칭.

package robots

import "testing"

func TestParseEmptyDisallowAddsNoRule(t *testing.T) {
	t.Run("empty disallow", func(t *testing.T) {
		rs := Parse([]byte("User-agent: *\nDisallow:\n"))
		if len(rs.Groups) != 1 {
			t.Fatalf("groups = %+v, want 1", rs.Groups)
		}
		if len(rs.Groups[0].Rules) != 0 {
			t.Errorf("expected no rules for empty Disallow, got %+v", rs.Groups[0].Rules)
		}
	})

	t.Run("empty allow (symmetric)", func(t *testing.T) {
		rs := Parse([]byte("User-agent: *\nAllow:\n"))
		if len(rs.Groups) != 1 {
			t.Fatalf("groups = %+v, want 1", rs.Groups)
		}
		if len(rs.Groups[0].Rules) != 0 {
			t.Errorf("expected no rules for empty Allow, got %+v", rs.Groups[0].Rules)
		}
	})

	t.Run("disallow slash is kept", func(t *testing.T) {
		rs := Parse([]byte("User-agent: *\nDisallow: /\n"))
		if len(rs.Groups) != 1 || len(rs.Groups[0].Rules) != 1 {
			t.Fatalf("groups = %+v, want one rule", rs.Groups)
		}
		if rs.Groups[0].Rules[0].Pattern != "/" || rs.Groups[0].Rules[0].Allow {
			t.Errorf("expected Disallow: / kept, got %+v", rs.Groups[0].Rules[0])
		}
	})
}
