//ff:func feature=extract type=helper control=sequence
//ff:what Apply가 PASS시 Extracted 채우고 본문 반환·State TODO 유지, Lang 미덮어쓰기, SKIP 잠금(구조화없음/본문부족), non-TODO 무변경을 검증한다.

package extract

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestApply(t *testing.T) {
	t.Run("pass fills Extracted and returns body, state stays TODO", func(t *testing.T) {
		in := []byte(`<html><head>` +
			`<script type="application/ld+json">{"@type":"NewsArticle","headline":"AH","author":{"name":"Auth"},"datePublished":"2026-01-01","articleBody":"` + longBody + `","inLanguage":"en"}</script>` +
			`</head><body></body></html>`)
		a := &session.Article{State: session.TODO}
		body, ok := Apply(a, in)
		if !ok {
			t.Fatalf("expected ok=true")
		}
		if body != longBody {
			t.Fatalf("body = %q", body)
		}
		if a.State != session.TODO {
			t.Fatalf("state = %v, want TODO (pass leaves TODO for phase006)", a.State)
		}
		if a.Extracted == nil {
			t.Fatalf("Extracted nil")
		}
		if a.Extracted.Title != "AH" || a.Extracted.Author != "Auth" ||
			a.Extracted.PublishedAt != "2026-01-01" || a.Extracted.Source != "jsonld" ||
			a.Extracted.BodyLen != len(longBody) {
			t.Fatalf("Extracted = %+v", a.Extracted)
		}
		if a.Lang != "en" {
			t.Fatalf("Lang = %q, want en", a.Lang)
		}
	})
	t.Run("does not overwrite pre-set Lang", func(t *testing.T) {
		in := []byte(`<html><head>` +
			`<script type="application/ld+json">{"@type":"Article","headline":"H","articleBody":"` + longBody + `","inLanguage":"de"}</script>` +
			`</head><body></body></html>`)
		a := &session.Article{State: session.TODO, Lang: "ko"}
		if _, ok := Apply(a, in); !ok {
			t.Fatalf("expected ok=true")
		}
		if a.Lang != "ko" {
			t.Fatalf("Lang overwritten: %q", a.Lang)
		}
	})
	t.Run("skip locks SKIPPED with reason", func(t *testing.T) {
		in := []byte(`<html><body>no structured data here</body></html>`)
		a := &session.Article{State: session.TODO}
		body, ok := Apply(a, in)
		if ok || body != "" {
			t.Fatalf("expected skip, got body=%q ok=%v", body, ok)
		}
		if a.State != session.SKIPPED {
			t.Fatalf("state = %v, want SKIPPED", a.State)
		}
		if a.SkipReason != SkipNoStructured {
			t.Fatalf("skipReason = %q", a.SkipReason)
		}
		if a.Extracted != nil {
			t.Fatalf("Extracted should be nil on skip")
		}
	})
	t.Run("skip body too short", func(t *testing.T) {
		in := []byte(`<html><head>` +
			`<script type="application/ld+json">{"@type":"Article","headline":"H","articleBody":"tiny"}</script>` +
			`</head><body>tiny</body></html>`)
		a := &session.Article{State: session.TODO}
		if _, ok := Apply(a, in); ok {
			t.Fatalf("expected skip")
		}
		if a.State != session.SKIPPED || a.SkipReason != SkipBodyTooShort {
			t.Fatalf("state=%v reason=%q", a.State, a.SkipReason)
		}
	})
	t.Run("non-TODO article untouched", func(t *testing.T) {
		a := &session.Article{State: session.PASS}
		body, ok := Apply(a, []byte(`<html></html>`))
		if ok || body != "" {
			t.Fatalf("expected ok=false body empty")
		}
		if a.State != session.PASS {
			t.Fatalf("state mutated: %v", a.State)
		}
		if a.Extracted != nil || a.SkipReason != "" {
			t.Fatalf("non-TODO mutated: %+v", a)
		}
	})
}
