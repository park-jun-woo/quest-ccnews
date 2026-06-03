//ff:type feature=session type=model
//ff:what 한 번의 수집 세션 전체 상태. user agent, ingestion 커서, 호스트 캐시, 기사 작업 목록을 담는다.

package session

// Session: the complete runtime state for one collection session (session.json).
type Session struct {
	Version   int              `json:"version"`
	UserAgent string           `json:"user_agent"`
	Ingestion Ingestion        `json:"ingestion"`
	Hosts     map[string]*Host `json:"hosts"`
	Articles  []*Article       `json:"articles"`
}
