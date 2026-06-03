//ff:func feature=extract type=helper control=sequence
//ff:what Parse가 jsonld(articleBody/폴백)·og 폴백·구조화없음 각 경로에서 Source와 BodyText를 올바르게 채우는지 검증한다.

package extract

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("jsonld with articleBody", func(t *testing.T) {
		in := []byte(`<html><head>` +
			`<script type="application/ld+json">{"@type":"NewsArticle","headline":"JH","articleBody":"` + longBody + `"}</script>` +
			`</head><body><p>ignored fallback</p></body></html>`)
		r := Parse(in)
		if r.Source != "jsonld" {
			t.Fatalf("source = %q, want jsonld", r.Source)
		}
		if r.Fields.Title != "JH" {
			t.Fatalf("title = %q", r.Fields.Title)
		}
		if r.BodyText != longBody {
			t.Fatalf("bodyText should be articleBody, got %q", r.BodyText)
		}
	})
	t.Run("jsonld without articleBody falls back to stripped HTML", func(t *testing.T) {
		in := []byte(`<html><head>` +
			`<script type="application/ld+json">{"@type":"Article","headline":"NoBody"}</script>` +
			`</head><body>Fallback text here</body></html>`)
		r := Parse(in)
		if r.Source != "jsonld" {
			t.Fatalf("source = %q", r.Source)
		}
		if !strings.Contains(r.BodyText, "Fallback text here") {
			t.Fatalf("bodyText = %q, want stripped fallback", r.BodyText)
		}
	})
	t.Run("og fallback when no jsonld article", func(t *testing.T) {
		in := []byte(`<html><head>` +
			`<meta property="og:title" content="OGT">` +
			`</head><body>Some body text for og</body></html>`)
		r := Parse(in)
		if r.Source != "og" {
			t.Fatalf("source = %q, want og", r.Source)
		}
		if r.Fields.Title != "OGT" {
			t.Fatalf("title = %q", r.Fields.Title)
		}
		if !strings.Contains(r.BodyText, "Some body text for og") {
			t.Fatalf("bodyText = %q", r.BodyText)
		}
	})
	t.Run("no structured data -> empty source, stripped body", func(t *testing.T) {
		in := []byte(`<html><body>Just plain content</body></html>`)
		r := Parse(in)
		if r.Source != "" {
			t.Fatalf("source = %q, want empty", r.Source)
		}
		if !strings.Contains(r.BodyText, "Just plain content") {
			t.Fatalf("bodyText = %q", r.BodyText)
		}
	})
}
