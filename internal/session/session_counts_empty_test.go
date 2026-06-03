//ff:func feature=article type=helper control=sequence
//ff:what Counts가 빈 세션에서 모든 상태를 0으로 집계하는지 검증한다.

package session

import (
	"reflect"
	"testing"
)

func TestCountsEmpty(t *testing.T) {
	s := &Session{Articles: nil}
	want := map[State]int{
		TODO: 0, PASS: 0, REVIEW: 0, DONE: 0, BLOCKED: 0, SKIPPED: 0,
	}
	if got := s.Counts(); !reflect.DeepEqual(got, want) {
		t.Errorf("Counts() = %v, want %v", got, want)
	}
}
