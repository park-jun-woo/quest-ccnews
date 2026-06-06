//ff:func feature=robots type=helper control=sequence
//ff:what putHostRobots 단위 테스트. fetch된 robots 레코드를 Session.Meta 호스트 캐시(metaHosts)에 적재하는지 검증한다. ① 빈 Meta에 새 host 추가. ② 기존 다른 host 보존(generic map 라운드트립). ③ 같은 host 재적재 시 Robots 갱신·기타 필드 보존. ④ Save/Load 라운드트립 후에도 host 키 보존.
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPutHostRobots(t *testing.T) {
	t.Run("adds a new host to empty Meta", func(t *testing.T) {
		s := quest.New()
		rec := &session.Robots{Status: "ok", RobotsURL: "https://a.com/robots.txt"}
		putHostRobots(s, "a.com", rec)

		hosts := metaHosts(t, s)
		if hosts["a.com"] == nil || hosts["a.com"].Robots == nil {
			t.Fatalf("a.com not recorded: %#v", hosts)
		}
		if hosts["a.com"].Robots.Status != "ok" {
			t.Errorf("Robots.Status = %q, want ok", hosts["a.com"].Robots.Status)
		}
	})

	t.Run("preserves other hosts already cached", func(t *testing.T) {
		s := quest.New()
		// Pre-seed Meta as the post-Load generic shape would deliver it: a typed map
		// here, then putHostRobots JSON round-trips it.
		s.SetMeta(session.MetaHosts, map[string]*session.Host{
			"old.com": {MediaName: "Old Media", Robots: &session.Robots{Status: "missing"}},
		})

		putHostRobots(s, "new.com", &session.Robots{Status: "ok"})

		hosts := metaHosts(t, s)
		if hosts["old.com"] == nil {
			t.Fatalf("old.com dropped: %#v", hosts)
		}
		if hosts["old.com"].MediaName != "Old Media" {
			t.Errorf("old.com MediaName lost: %q", hosts["old.com"].MediaName)
		}
		if hosts["new.com"] == nil {
			t.Fatalf("new.com not added: %#v", hosts)
		}
	})

	t.Run("re-recording same host updates Robots but keeps other fields", func(t *testing.T) {
		s := quest.New()
		s.SetMeta(session.MetaHosts, map[string]*session.Host{
			"h.com": {MediaName: "Keep Me", Robots: &session.Robots{Status: "missing"}},
		})

		putHostRobots(s, "h.com", &session.Robots{Status: "unreachable"})

		hosts := metaHosts(t, s)
		if hosts["h.com"].MediaName != "Keep Me" {
			t.Errorf("MediaName clobbered: %q", hosts["h.com"].MediaName)
		}
		if hosts["h.com"].Robots.Status != "unreachable" {
			t.Errorf("Robots not updated: %q", hosts["h.com"].Robots.Status)
		}
	})

	t.Run("survives Save/Load round-trip", func(t *testing.T) {
		s := quest.New()
		putHostRobots(s, "rt.com", &session.Robots{Status: "ok"})

		path := t.TempDir() + "/session.json"
		if err := s.Save(path); err != nil {
			t.Fatal(err)
		}
		reloaded, err := quest.Load(path)
		if err != nil {
			t.Fatal(err)
		}
		v, ok := reloaded.GetMeta(session.MetaHosts)
		if !ok {
			t.Fatal("Meta[hosts] absent after round-trip")
		}
		// After Load a Meta value is a generic map[string]any.
		m, ok := v.(map[string]any)
		if !ok || m["rt.com"] == nil {
			t.Fatalf("rt.com missing after round-trip: %#v", v)
		}

		// putHostRobots tolerates the generic post-Load shape and still preserves it.
		putHostRobots(reloaded, "rt2.com", &session.Robots{Status: "ok"})
		hosts := metaHosts(t, reloaded)
		if hosts["rt.com"] == nil || hosts["rt2.com"] == nil {
			t.Fatalf("post-Load update lost a host: %#v", hosts)
		}
	})
}
