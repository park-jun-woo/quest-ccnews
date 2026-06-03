//ff:func feature=article type=helper control=iteration dimension=1
//ff:what 다음에 처리할 TODO 기사를 반환한다. 잠긴 상태는 절대 다시 집지 않는다(래칫).

package session

// NextTODO: the next article to process. Locked states
// (PASS/REVIEW/DONE/BLOCKED/SKIPPED) are never picked again (ratchet).
func (s *Session) NextTODO() *Article {
	for _, a := range s.Articles {
		if a.State == TODO {
			return a
		}
	}
	return nil
}
