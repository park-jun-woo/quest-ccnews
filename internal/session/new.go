//ff:func feature=session type=helper control=sequence
//ff:what 주어진 user agent와 ingestion source로 빈 세션을 새로 만든다(호스트 캐시·기사 목록 비움).

package session

// New: create a fresh session for the given user agent and ingestion source.
// Starts with an empty host cache and an empty article work-list.
func New(userAgent, source string) *Session {
	return &Session{
		Version:   1,
		UserAgent: userAgent,
		Ingestion: Ingestion{
			Source:         source,
			ProcessedWarcs: []string{},
		},
		Hosts:    map[string]*Host{},
		Articles: []*Article{},
	}
}
