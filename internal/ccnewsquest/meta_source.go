//ff:func feature=gate type=helper control=sequence
//ff:what Session.Meta에서 WARC 재독 설정(cacheDir·userAgent)을 소싱한다(Phase013 B). Meta의 cache_dir·user_agent 문자열이 있으면 그것을, 없으면 ccnewsDef 리시버 기본값을 쓴다(하위호환: 미기록 세션은 기존 동작). run이 cacheDir를 절대경로로 정규화해 기록하므로 CWD 의존이 사라진다.
package ccnewsquest

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// sourceConfig resolves the WARC re-read configuration (cache dir + crawl UA) from
// the session Meta slot, falling back to the ccnewsDef receiver defaults when a key
// is absent (Phase013 B). The `run` command writes an absolute cache_dir into Meta,
// so submit/next re-read from that path regardless of CWD; older sessions without
// the keys keep the receiver's compile-time defaults (backward compatible).
func (d ccnewsDef) sourceConfig(s *quest.Session) (userAgent, cacheDir string) {
	userAgent, cacheDir = d.userAgent, d.cacheDir
	if s == nil {
		return userAgent, cacheDir
	}
	if v, ok := s.GetMeta(session.MetaUserAgent); ok {
		if str, ok := v.(string); ok && str != "" {
			userAgent = str
		}
	}
	if v, ok := s.GetMeta(session.MetaCacheDir); ok {
		if str, ok := v.(string); ok && str != "" {
			cacheDir = str
		}
	}
	return userAgent, cacheDir
}
