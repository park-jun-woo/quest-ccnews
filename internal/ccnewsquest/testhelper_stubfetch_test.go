//ff:func feature=robots type=helper control=sequence
//ff:what robots 테스트 공용 헬퍼 stubFetch. robotsFetch를 rec/rs를 반환하며 호출 횟수를 세는 스텁으로 교체하고 t.Cleanup으로 원복한다. 호출 카운터 포인터를 반환해 fetch 1회/캐시 히트 검증에 쓴다.
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// stubFetch swaps robotsFetch for one returning rec/rs and counting calls; it
// restores the original on cleanup.
func stubFetch(t *testing.T, rec *session.Robots, rs *robots.Ruleset) *int {
	t.Helper()
	calls := 0
	orig := robotsFetch
	robotsFetch = func(string, string) (*session.Robots, *robots.Ruleset, error) {
		calls++
		return rec, rs, nil
	}
	t.Cleanup(func() { robotsFetch = orig })
	return &calls
}
