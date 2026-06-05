//ff:func feature=ingestion type=helper control=sequence
//ff:what decodeMeta 단위테스트. 키 부재면 (false,nil)이고 dst 무수정, 존재하면 JSON 왕복으로 디코드해 (true,nil), 타입 불일치(언마샬 실패)·마샬 불가(channel) 입력은 (true,err)를 돌려주는지 검증한다.

package runcmd

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestDecodeMeta(t *testing.T) {
	t.Run("absent key returns ok=false without touching dst", func(t *testing.T) {
		s := quest.New()
		var dst session.Ingestion
		ok, err := decodeMeta(s, "missing", &dst)
		if ok || err != nil {
			t.Fatalf("ok=%v err=%v, want false,nil", ok, err)
		}
		if dst.Source != "" {
			t.Errorf("dst mutated: %+v", dst)
		}
	})

	t.Run("present value decodes via JSON round-trip", func(t *testing.T) {
		s := quest.New()
		// Stored as a generic map (mirrors a post-Load Meta value).
		s.SetMeta("ingestion", map[string]any{
			"source":          "cc-news",
			"processed_warcs": []any{"w1", "w2"},
		})
		var dst session.Ingestion
		ok, err := decodeMeta(s, "ingestion", &dst)
		if !ok || err != nil {
			t.Fatalf("ok=%v err=%v, want true,nil", ok, err)
		}
		if dst.Source != "cc-news" || len(dst.ProcessedWarcs) != 2 {
			t.Errorf("decoded = %+v", dst)
		}
	})

	t.Run("unmarshal type mismatch returns ok=true with error", func(t *testing.T) {
		s := quest.New()
		s.SetMeta("ingestion", "not-an-object") // string → struct fails
		var dst session.Ingestion
		ok, err := decodeMeta(s, "ingestion", &dst)
		if !ok {
			t.Errorf("ok = false, want true (key present)")
		}
		if err == nil {
			t.Errorf("err = nil, want unmarshal error")
		}
	})

	t.Run("marshal error returns ok=true with error", func(t *testing.T) {
		s := quest.New()
		// channels are not JSON-marshalable → json.Marshal fails.
		s.SetMeta("bad", make(chan int))
		var dst session.Ingestion
		ok, err := decodeMeta(s, "bad", &dst)
		if !ok {
			t.Errorf("ok = false, want true (key present)")
		}
		if err == nil {
			t.Errorf("err = nil, want marshal error")
		}
	})
}
