//ff:func feature=robots type=helper control=sequence
//ff:what Phase013 A 검증 게이트(pick-time robots). robotsFetch를 카운터 스텁으로 주입해 검증한다. ① 같은 host 기사 2건 submit → fetch 정확히 1회(2번째는 in-memory ruleset 캐시 히트). ② 거부 host 기사 submit → Prepare가 OutBlock short verdict 반환(BLOCKED, WARC 재독 없이 단락). ③ fetch 결과 session.Robots 레코드가 Session.Meta 호스트 캐시에 적재되고 저장/복원 라운드트립으로 보존. (시드 중 fetch 0은 bridge가 robots를 전혀 호출하지 않음으로 보장 — runcmd TestBridge.)
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareRobotsPickTime(t *testing.T) {
	t.Run("same host twice → robots fetched exactly once", func(t *testing.T) {
		cacheDir, file := writeWarcHTML(t, passHTML)
		d := Def("ua", cacheDir)                                                // fresh shared robots cache
		calls := stubFetch(t, &session.Robots{Status: "ok"}, &robots.Ruleset{}) // empty → allow
		s := quest.New()

		for i := 0; i < 2; i++ {
			it := &quest.Item{Key: "https://h.com/a"}
			if err := it.SetPayload(&session.Article{
				URL: "https://h.com/a", Host: "h.com", State: session.TODO,
				WARC: &session.WARCLoc{File: file, Offset: 0},
			}); err != nil {
				t.Fatal(err)
			}
			if _, v, err := d.Prepare(s, it, []byte(`{"who":{"value":"x","anchors":["article body"]}}`)); err != nil || v != nil {
				t.Fatalf("Prepare #%d: v=%+v err=%v", i, v, err)
			}
		}
		if *calls != 1 {
			t.Errorf("robots fetch calls = %d, want exactly 1 (second is a cache hit)", *calls)
		}
	})

	t.Run("denied host → OutBlock short verdict (BLOCKED)", func(t *testing.T) {
		d := Def("ua", t.TempDir())
		stubFetch(t, &session.Robots{Status: "ok"}, denyRuleset())
		s := quest.New()

		it := &quest.Item{Key: "https://deny.com/x"}
		if err := it.SetPayload(&session.Article{
			URL: "https://deny.com/x", Host: "deny.com", State: session.TODO,
			// A bogus WARC: if robots did NOT short-circuit, ReadBody would error,
			// proving the block happens before (and without) the body re-read.
			WARC: &session.WARCLoc{File: "nope.warc", Offset: 0},
		}); err != nil {
			t.Fatal(err)
		}
		_, v, err := d.Prepare(s, it, []byte(`{"who":{"value":"x","anchors":["y"]}}`))
		if err != nil {
			t.Fatalf("Prepare denied host err = %v (should short-circuit, not read WARC)", err)
		}
		if v == nil || v.Outcome != quest.OutBlock {
			t.Fatalf("verdict = %+v, want OutBlock", v)
		}
		var a2 session.Article
		if err := it.DecodePayload(&a2); err != nil {
			t.Fatal(err)
		}
		if a2.State != session.BLOCKED || a2.SkipReason == "" {
			t.Errorf("payload State=%q reason=%q, want BLOCKED + robots reason", a2.State, a2.SkipReason)
		}
	})

	t.Run("robots record cached into Meta survives Save/Load round-trip", func(t *testing.T) {
		cacheDir, file := writeWarcHTML(t, passHTML)
		d := Def("ua", cacheDir)
		stubFetch(t, &session.Robots{Status: "ok", RobotsURL: "https://h.com/robots.txt"}, &robots.Ruleset{})
		s := quest.New()

		it := &quest.Item{Key: "https://h.com/a"}
		if err := it.SetPayload(&session.Article{
			URL: "https://h.com/a", Host: "h.com", State: session.TODO,
			WARC: &session.WARCLoc{File: file, Offset: 0},
		}); err != nil {
			t.Fatal(err)
		}
		if _, v, err := d.Prepare(s, it, []byte(`{"who":{"value":"x","anchors":["article body"]}}`)); err != nil || v != nil {
			t.Fatalf("Prepare: v=%+v err=%v", v, err)
		}

		// Round-trip the session through disk like submit's Save + a later Load.
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
			t.Fatalf("Meta[hosts] absent after round-trip")
		}
		// After Load it is a generic map; assert the host key is present.
		m, ok := v.(map[string]any)
		if !ok || m["h.com"] == nil {
			t.Fatalf("Meta[hosts][h.com] missing after round-trip: %#v", v)
		}
	})
}
