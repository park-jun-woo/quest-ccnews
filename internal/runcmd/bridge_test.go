//ff:func feature=ingestion type=helper control=sequence
//ff:what bridge 단위테스트(Phase013 A: robots 시드-fetch 없음). 새 기사를 URL키·payload·TODO Item으로 시드하고 처리 뒤 scratch.Articles를 비우는지, nil·빈URL·이미시드·배치내중복을 건너뛰는지, 커서/hosts/UA/cache_dir를 Meta로 보존하는지 검증한다. robots 판정은 pick-time(Prepare)로 이동했으므로 시드는 모든 기사를 TODO로 둔다.

package runcmd

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestBridge(t *testing.T) {
	t.Run("seeds new articles as TODO with URL key and payload", func(t *testing.T) {
		scratch := session.New("ua", "cc-news")
		a := &session.Article{URL: "https://ex.com/a", Host: "ex.com"}
		scratch.Articles = []*session.Article{a}
		s := quest.New()

		n := bridge(scratch, s, "")

		if n != 1 {
			t.Fatalf("seeded = %d, want 1", n)
		}
		if len(s.Items) != 1 {
			t.Fatalf("items = %d, want 1", len(s.Items))
		}
		it := s.Items[0]
		if it.Key != a.URL {
			t.Errorf("Key = %q, want %q", it.Key, a.URL)
		}
		if it.State != quest.TODO {
			t.Errorf("State = %q, want TODO", it.State)
		}
		var got session.Article
		if err := it.DecodePayload(&got); err != nil {
			t.Fatalf("DecodePayload: %v", err)
		}
		if got.URL != a.URL || got.Host != a.Host {
			t.Errorf("payload = %+v, want article %+v", got, *a)
		}
		// scratch.Articles truncated after fold.
		if len(scratch.Articles) != 0 {
			t.Errorf("scratch.Articles len = %d, want 0", len(scratch.Articles))
		}
	})

	t.Run("skips nil, empty-URL, and already-seeded articles", func(t *testing.T) {
		s := quest.New()
		s.Items = []*quest.Item{{Key: "https://ex.com/dup", State: quest.TODO}}
		scratch := session.New("ua", "cc-news")
		scratch.Articles = []*session.Article{
			nil,
			{URL: ""},
			{URL: "https://ex.com/dup", Host: "ex.com"}, // already seeded
			{URL: "https://ex.com/dup", Host: "ex.com"}, // dup within this batch
			{URL: "https://ex.com/new", Host: "ex.com"},
		}

		n := bridge(scratch, s, "")

		if n != 1 {
			t.Fatalf("seeded = %d, want 1 (only the new URL)", n)
		}
		if len(s.Items) != 2 {
			t.Fatalf("items = %d, want 2", len(s.Items))
		}
		if s.Items[1].Key != "https://ex.com/new" {
			t.Errorf("new item key = %q", s.Items[1].Key)
		}
	})

	t.Run("preserves cursor/hosts/UA/cache_dir into Meta", func(t *testing.T) {
		scratch := session.New("crawl-ua", "cc-news")
		scratch.Ingestion.ProcessedWarcs = []string{"w1.warc.gz"}
		scratch.Hosts = map[string]*session.Host{"ex.com": {MediaName: "Ex"}}
		s := quest.New()

		bridge(scratch, s, "/abs/warc-cache")

		if v, ok := s.GetMeta(metaUserAgent); !ok || v.(string) != "crawl-ua" {
			t.Errorf("Meta[user_agent] = %v, ok=%v", v, ok)
		}
		ing, ok := s.GetMeta(metaIngestion)
		if !ok || ing.(session.Ingestion).Source != "cc-news" {
			t.Errorf("Meta[ingestion] = %v, ok=%v", ing, ok)
		}
		hosts, ok := s.GetMeta(metaHosts)
		if !ok || hosts.(map[string]*session.Host)["ex.com"] == nil {
			t.Errorf("Meta[hosts] = %v, ok=%v", hosts, ok)
		}
		if v, ok := s.GetMeta(metaCacheDir); !ok || v.(string) != "/abs/warc-cache" {
			t.Errorf("Meta[cache_dir] = %v, ok=%v, want /abs/warc-cache", v, ok)
		}
	})
}
