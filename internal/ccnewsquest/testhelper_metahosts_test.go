//ff:func feature=host type=helper control=sequence
//ff:what 호스트 캐시 테스트 공용 헬퍼 metaHosts. 코드 본체와 동일한 방식(generic post-Load shape를 JSON 라운드트립)으로 Session.Meta에서 타입드 호스트 캐시(map[string]*session.Host)를 복원해 반환한다.
package ccnewsquest

import (
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// metaHosts decodes the typed host cache back out of Meta the same way the code
// under test does (JSON round-trip through the generic post-Load shape).
func metaHosts(t *testing.T, s *quest.Session) map[string]*session.Host {
	t.Helper()
	v, ok := s.GetMeta(session.MetaHosts)
	if !ok {
		t.Fatal("Meta[hosts] absent")
	}
	hosts := map[string]*session.Host{}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal Meta[hosts]: %v", err)
	}
	if err := json.Unmarshal(b, &hosts); err != nil {
		t.Fatalf("unmarshal Meta[hosts]: %v", err)
	}
	return hosts
}
