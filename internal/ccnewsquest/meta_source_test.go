//ff:func feature=gate type=helper control=sequence
//ff:what sourceConfig 단위 테스트. Meta의 cache_dir·user_agent가 있으면 그것을, 없으면 리시버 기본값을 소싱하는지 검증한다. ① nil 세션 → 리시버 기본값. ② Meta 미설정 → 리시버 기본값(하위호환). ③ Meta 값 존재 → Meta 우선. ④ Meta가 빈 문자열·비문자열이면 무시하고 리시버 기본값 유지.
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestSourceConfig(t *testing.T) {
	d := ccnewsDef{userAgent: "ua-default", cacheDir: "/cache-default"}

	t.Run("nil session → receiver defaults", func(t *testing.T) {
		ua, cd := d.sourceConfig(nil)
		if ua != "ua-default" || cd != "/cache-default" {
			t.Fatalf("got (%q,%q), want receiver defaults", ua, cd)
		}
	})

	t.Run("no Meta keys → receiver defaults (backward compatible)", func(t *testing.T) {
		ua, cd := d.sourceConfig(quest.New())
		if ua != "ua-default" || cd != "/cache-default" {
			t.Fatalf("got (%q,%q), want receiver defaults", ua, cd)
		}
	})

	t.Run("Meta values override receiver defaults", func(t *testing.T) {
		s := quest.New()
		s.SetMeta(session.MetaUserAgent, "ua-meta")
		s.SetMeta(session.MetaCacheDir, "/cache-meta")
		ua, cd := d.sourceConfig(s)
		if ua != "ua-meta" || cd != "/cache-meta" {
			t.Fatalf("got (%q,%q), want Meta values", ua, cd)
		}
	})

	t.Run("empty-string Meta values are ignored", func(t *testing.T) {
		s := quest.New()
		s.SetMeta(session.MetaUserAgent, "")
		s.SetMeta(session.MetaCacheDir, "")
		ua, cd := d.sourceConfig(s)
		if ua != "ua-default" || cd != "/cache-default" {
			t.Fatalf("got (%q,%q), want receiver defaults (empty ignored)", ua, cd)
		}
	})

	t.Run("non-string Meta values are ignored", func(t *testing.T) {
		s := quest.New()
		s.SetMeta(session.MetaUserAgent, 42)
		s.SetMeta(session.MetaCacheDir, []string{"x"})
		ua, cd := d.sourceConfig(s)
		if ua != "ua-default" || cd != "/cache-default" {
			t.Fatalf("got (%q,%q), want receiver defaults (non-string ignored)", ua, cd)
		}
	})

	t.Run("only one Meta key set → other falls back", func(t *testing.T) {
		s := quest.New()
		s.SetMeta(session.MetaCacheDir, "/cache-meta")
		ua, cd := d.sourceConfig(s)
		if ua != "ua-default" || cd != "/cache-meta" {
			t.Fatalf("got (%q,%q), want (ua-default,/cache-meta)", ua, cd)
		}
	})
}
