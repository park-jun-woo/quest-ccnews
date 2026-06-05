//ff:func feature=gate type=helper control=sequence
//ff:what Def().Seed() 스모크. URL 목록을 TODO 아이템(Payload=*Article)으로 시드하고 공백 URL은 떨어뜨리며 Key·State·payload URL이 옳은지 검증한다.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestSeed(t *testing.T) {
	items, err := Def("ua", "cache").Seed([]string{"https://x/a", "  ", "https://x/b"})
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 2 {
		t.Fatalf("seeded %d items, want 2 (blank dropped)", len(items))
	}
	if items[0].Key != "https://x/a" || items[0].State != quest.TODO {
		t.Fatalf("item0 = %+v", items[0])
	}
	var a session.Article
	if err := items[0].DecodePayload(&a); err != nil || a.URL != "https://x/a" {
		t.Fatalf("item0 payload = %q (err %v)", string(items[0].Payload), err)
	}
}
