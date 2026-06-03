//ff:func feature=article type=helper control=sequence
//ff:what Counts가 여러 상태가 섞인 세션을 상태별로 정확히 집계하는지 검증한다.

package session

import (
	"reflect"
	"testing"
)

func TestCountsMixed(t *testing.T) {
	s := &Session{Articles: []*Article{
		{State: TODO},
		{State: TODO},
		{State: PASS},
		{State: REVIEW},
		{State: DONE},
		{State: BLOCKED},
		{State: SKIPPED},
	}}
	want := map[State]int{
		TODO: 2, PASS: 1, REVIEW: 1, DONE: 1, BLOCKED: 1, SKIPPED: 1,
	}
	if got := s.Counts(); !reflect.DeepEqual(got, want) {
		t.Errorf("Counts() = %v, want %v", got, want)
	}
}
