//ff:func feature=gate type=helper control=sequence
//ff:what skipVerdict가 Article을 payload에 되쓰고(SkipReason 보존) OutSkip verdict(structured-trust Fact, Actual=SkipReason)를 돌려주는지 검증한다.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestSkipVerdict(t *testing.T) {
	it := &quest.Item{Key: "https://example.com/a"}
	a := &session.Article{
		URL:        "https://example.com/a",
		State:      session.TODO,
		SkipReason: "no-structured",
	}

	v := skipVerdict(it, a)
	if v == nil || v.Outcome != quest.OutSkip {
		t.Fatalf("verdict = %+v, want OutSkip", v)
	}
	if len(v.Facts) != 1 {
		t.Fatalf("facts = %d, want 1", len(v.Facts))
	}
	f := v.Facts[0]
	if f.Rule != "structured-trust" || f.Where != "body" || f.Actual != "no-structured" {
		t.Fatalf("fact = %+v, want structured-trust/body/no-structured", f)
	}

	var a2 session.Article
	if err := it.DecodePayload(&a2); err != nil {
		t.Fatalf("DecodePayload: %v", err)
	}
	if a2.SkipReason != "no-structured" {
		t.Fatalf("payload SkipReason = %q, want %q", a2.SkipReason, "no-structured")
	}
}
