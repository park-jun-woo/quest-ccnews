//ff:func feature=gate type=helper control=sequence
//ff:what Phase013 B 검증 게이트. ① Meta에 절대 cache_dir가 있으면 Prepare가 리시버 cacheDir이 아닌 Meta 경로에서 WARC를 재독한다(CWD 의존 제거). ② cache_dir 미기록 세션은 리시버 기본값으로 동작(하위호환). ③ Meta user_agent도 같은 방식으로 소싱.
package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareSourcesCacheDirFromMeta(t *testing.T) {
	cacheDir, file := writeWarcHTML(t, passHTML)

	newItem := func() *quest.Item {
		it := &quest.Item{Key: "https://example.com/a"}
		if err := it.SetPayload(&session.Article{
			URL:   "https://example.com/a",
			State: session.TODO,
			WARC:  &session.WARCLoc{File: file, Offset: 0},
		}); err != nil {
			t.Fatal(err)
		}
		return it
	}

	t.Run("Meta cache_dir overrides a wrong receiver cacheDir", func(t *testing.T) {
		// Receiver cacheDir is bogus; only the absolute Meta cache_dir locates the WARC.
		d := Def("ua", "/nonexistent-receiver-cache")
		s := quest.New()
		s.SetMeta(session.MetaCacheDir, cacheDir)

		_, v, err := d.Prepare(s, newItem(), []byte(`{"who":{"value":"x","anchors":["article body"]}}`))
		if err != nil {
			t.Fatalf("Prepare with Meta cache_dir: %v (Meta path not used?)", err)
		}
		if v != nil {
			t.Fatalf("unexpected short verdict %+v", v)
		}
	})

	t.Run("no Meta cache_dir → receiver default (backward compatible)", func(t *testing.T) {
		d := Def("ua", cacheDir) // receiver holds the real dir
		s := quest.New()         // no Meta keys

		_, v, err := d.Prepare(s, newItem(), []byte(`{"who":{"value":"x","anchors":["article body"]}}`))
		if err != nil {
			t.Fatalf("Prepare fallback to receiver cacheDir: %v", err)
		}
		if v != nil {
			t.Fatalf("unexpected short verdict %+v", v)
		}
	})

	t.Run("wrong receiver + no Meta → error (proves Meta is what fixes it)", func(t *testing.T) {
		d := Def("ua", "/nonexistent-receiver-cache")
		_, _, err := d.Prepare(quest.New(), newItem(), []byte(`{"who":{"value":"x","anchors":["article body"]}}`))
		if err == nil {
			t.Fatalf("want ReadBody error when neither receiver nor Meta points at the WARC")
		}
	})
}
