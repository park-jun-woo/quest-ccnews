//ff:func feature=article type=helper control=iteration dimension=1
//ff:what 기사를 상태별로 집계한다(status 출력용).

package session

// Counts: tally articles by state (for status output).
func (s *Session) Counts() map[State]int {
	m := map[State]int{
		TODO:    0,
		PASS:    0,
		REVIEW:  0,
		DONE:    0,
		BLOCKED: 0,
		SKIPPED: 0,
	}
	for _, a := range s.Articles {
		m[a.State]++
	}
	return m
}
