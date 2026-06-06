//ff:func feature=gate type=helper control=iteration dimension=1
//ff:what lastFailureReason 표 테스트. 빈 로그→"", 비어있지 않으면 마지막 항목 Reason 반환.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestLastFailureReason(t *testing.T) {
	cases := []struct {
		name string
		log  []quest.Attempt
		want string
	}{
		{"empty log", nil, ""},
		{"single", []quest.Attempt{{Reason: "r1"}}, "r1"},
		{"last wins", []quest.Attempt{{Reason: "r1"}, {Reason: "r2"}}, "r2"},
		{"empty reason tail", []quest.Attempt{{Reason: "r1"}, {Reason: ""}}, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			it := &quest.Item{Log: c.log}
			if got := lastFailureReason(it); got != c.want {
				t.Fatalf("lastFailureReason = %q, want %q", got, c.want)
			}
		})
	}
}
