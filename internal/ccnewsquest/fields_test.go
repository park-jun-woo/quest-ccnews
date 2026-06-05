//ff:func feature=gate type=helper control=sequence
//ff:what event6Of 단위테스트. *session.Event6면 ok=true에 필수(who/what)·선택(when/where/how/why) 순서대로 namedField 리스트 반환. 비-Event6/nil 포인터면 ok=false. 라벨·포인터 매핑 검증.

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/gate"
)

func TestEvent6Of(t *testing.T) {
	t.Run("event6 unwraps with ordered field lists", func(t *testing.T) {
		ev := &session.Event6{
			Who: fld("Alice", "Alice"), What: fld("signed", "sign the treaty"),
			When: fld("Monday", "Monday"), Where: fld("Paris", "Paris"),
			How: fld("met", "met"), Why: fld("peace", "peace mattered"),
		}
		required, optional, ok := event6Of(gate.Context{Submission: ev})
		if !ok {
			t.Fatal("want ok=true for *session.Event6")
		}
		wantReq := []struct {
			name string
			f    *session.Field
		}{{"who", ev.Who}, {"what", ev.What}}
		if len(required) != len(wantReq) {
			t.Fatalf("required len = %d, want %d", len(required), len(wantReq))
		}
		for i, w := range wantReq {
			if required[i].name != w.name || required[i].f != w.f {
				t.Fatalf("required[%d] = {%q,%p}, want {%q,%p}", i, required[i].name, required[i].f, w.name, w.f)
			}
		}
		wantOpt := []struct {
			name string
			f    *session.Field
		}{{"when", ev.When}, {"where", ev.Where}, {"how", ev.How}, {"why", ev.Why}}
		if len(optional) != len(wantOpt) {
			t.Fatalf("optional len = %d, want %d", len(optional), len(wantOpt))
		}
		for i, w := range wantOpt {
			if optional[i].name != w.name || optional[i].f != w.f {
				t.Fatalf("optional[%d] = {%q,%p}, want {%q,%p}", i, optional[i].name, optional[i].f, w.name, w.f)
			}
		}
	})
	t.Run("non-event6 submission -> ok=false", func(t *testing.T) {
		req, opt, ok := event6Of(gate.Context{Submission: "nope"})
		if ok || req != nil || opt != nil {
			t.Fatalf("want ok=false,nil,nil for non-Event6; got %v,%v,%v", ok, req, opt)
		}
	})
	t.Run("nil *Event6 -> ok=false", func(t *testing.T) {
		var ev *session.Event6
		_, _, ok := event6Of(gate.Context{Submission: ev})
		if ok {
			t.Fatal("want ok=false for typed-nil *session.Event6")
		}
	})
}
