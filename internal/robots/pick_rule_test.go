//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what pickRule이 비매칭은 best 유지, 더 긴 패턴은 승, 동률이면 Allow가 Disallow를 이기는 RFC 9309 최장 일치 우선을 테이블로 검증한다.

package robots

import "testing"

func TestPickRule(t *testing.T) {
	disShort := &Rule{Allow: false, Pattern: "/a"}     // len 2
	disLong := &Rule{Allow: false, Pattern: "/a/long"} // len 7
	allowEq := &Rule{Allow: true, Pattern: "/x"}       // len 2
	allowA := &Rule{Allow: true, Pattern: "/a"}        // len 2
	noMatch := &Rule{Allow: false, Pattern: "/nope"}

	tests := []struct {
		name     string
		r        *Rule
		norm     string
		best     *Rule
		bestLen  int
		wantBest *Rule
		wantLen  int
	}{
		{"non-match keeps best", noMatch, "/a/long/page", disShort, 2, disShort, 2},
		{"longer wins from nil", disLong, "/a/long/page", nil, -1, disLong, 7},
		{"shorter loses to best", disShort, "/a/long/page", disLong, 7, disLong, 7},
		{"tie allow beats disallow", allowEq, "/x", disShort, 2, allowEq, 2},
		{"tie allow over nil best", allowEq, "/x", nil, 2, allowEq, 2},
		{"tie disallow does not displace allow", disShort, "/a", allowA, 2, allowA, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBest, gotLen := pickRule(tt.r, tt.norm, tt.best, tt.bestLen)
			if gotBest != tt.wantBest || gotLen != tt.wantLen {
				t.Errorf("pickRule = (%+v, %d), want (%+v, %d)", gotBest, gotLen, tt.wantBest, tt.wantLen)
			}
		})
	}
}
