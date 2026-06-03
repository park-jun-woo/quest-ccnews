//ff:func feature=extract type=helper control=sequence
//ff:what Gate가 PASS·구조화없음(빈source/제목없음)·본문부족·MinBodyLen 경계(정확히/하나아래)를 올바르게 판정하는지 검증한다.

package extract

import (
	"strings"
	"testing"
)

func TestGate(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		r := Result{Fields: Fields{Title: "T"}, Source: "jsonld", BodyText: longBody}
		d := Gate(r)
		if !d.Pass || d.SkipReason != "" {
			t.Fatalf("expected pass, got %+v", d)
		}
		if d.Source != "jsonld" || d.BodyLen != len(longBody) {
			t.Fatalf("metadata wrong: %+v", d)
		}
	})
	t.Run("skip no structured (empty source)", func(t *testing.T) {
		d := Gate(Result{Fields: Fields{Title: "T"}, Source: "", BodyText: longBody})
		if d.Pass || d.SkipReason != SkipNoStructured {
			t.Fatalf("got %+v", d)
		}
	})
	t.Run("skip no structured (no title)", func(t *testing.T) {
		d := Gate(Result{Fields: Fields{}, Source: "jsonld", BodyText: longBody})
		if d.Pass || d.SkipReason != SkipNoStructured {
			t.Fatalf("got %+v", d)
		}
	})
	t.Run("skip body too short", func(t *testing.T) {
		d := Gate(Result{Fields: Fields{Title: "T"}, Source: "jsonld", BodyText: "short"})
		if d.Pass || d.SkipReason != SkipBodyTooShort {
			t.Fatalf("got %+v", d)
		}
	})
	t.Run("boundary exactly MinBodyLen passes", func(t *testing.T) {
		body := strings.Repeat("a", MinBodyLen)
		d := Gate(Result{Fields: Fields{Title: "T"}, Source: "og", BodyText: body})
		if !d.Pass {
			t.Fatalf("exactly MinBodyLen should pass, got %+v", d)
		}
	})
	t.Run("boundary one below MinBodyLen skips", func(t *testing.T) {
		body := strings.Repeat("a", MinBodyLen-1)
		d := Gate(Result{Fields: Fields{Title: "T"}, Source: "og", BodyText: body})
		if d.Pass || d.SkipReason != SkipBodyTooShort {
			t.Fatalf("one below should skip, got %+v", d)
		}
	})
}
