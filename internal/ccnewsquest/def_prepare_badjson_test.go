//ff:func feature=gate type=helper control=sequence
//ff:what Def().Prepare() 잘못된 JSON 가드(Phase015 A 반전). 디코드 불가능한 입력에서 Go 에러가 아니라 OutFail short verdict(RootCause=="event6-json", err==nil)를 돌려주는지 검증한다 — 한 번의 포맷 잡음이 무인 루프를 죽이지 않게.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareBadJSON(t *testing.T) {
	it := &quest.Item{Key: "https://x/a"}
	if err := it.SetPayload(&session.Article{URL: "https://x/a"}); err != nil {
		t.Fatal(err)
	}
	_, v, err := Def("ua", "cache").Prepare(quest.New(), it, []byte("{not json"))
	if err != nil {
		t.Fatalf("err = %v, want nil (format noise is a retryable FAIL, not a Go error)", err)
	}
	if v == nil || v.Outcome != quest.OutFail {
		t.Fatalf("verdict = %+v, want OutFail short verdict", v)
	}
	if v.RootCause != "event6-json" {
		t.Fatalf("RootCause = %q, want %q (so rule coaching fires)", v.RootCause, "event6-json")
	}
}
