//ff:func feature=output type=helper control=sequence
//ff:what renderAudit(BLOCKED/SKIPPED)가 url/host/status/reason과 crawl_allowed만 담고 수집필드는 누출하지 않으며, host/robots 유무를 올바르게 처리하는지 검증한다.

package output

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRenderAudit(t *testing.T) {
	t.Run("blocked with host", func(t *testing.T) {
		a := &session.Article{
			URL: "u", Host: "h", State: session.BLOCKED, SkipReason: "robots refused",
		}
		host := &session.Host{Robots: &session.Robots{CrawlAllowed: false}}
		r := Render(a, host)
		if r == nil {
			t.Fatal("Render() = nil")
		}
		if r.Status != "BLOCKED" || r.Reason != "robots refused" {
			t.Errorf("audit fields wrong: %+v", r)
		}
		if r.CrawlAllowed == nil || *r.CrawlAllowed != false {
			t.Errorf("CrawlAllowed = %v, want *false", r.CrawlAllowed)
		}
		if r.AnchorSummary != "" || r.Event6 != nil || r.MediaName != "" {
			t.Errorf("audit record leaked collection fields: %+v", r)
		}
	})

	t.Run("skipped no host", func(t *testing.T) {
		a := &session.Article{URL: "u", Host: "h", State: session.SKIPPED, SkipReason: "no structured data"}
		r := Render(a, nil)
		if r.Status != "SKIPPED" || r.Reason != "no structured data" {
			t.Errorf("audit fields wrong: %+v", r)
		}
		if r.CrawlAllowed != nil {
			t.Errorf("CrawlAllowed = %v, want nil (no host)", r.CrawlAllowed)
		}
	})

	t.Run("host no robots", func(t *testing.T) {
		a := &session.Article{State: session.BLOCKED}
		host := &session.Host{} // Robots nil
		r := Render(a, host)
		if r.CrawlAllowed != nil {
			t.Errorf("CrawlAllowed = %v, want nil (no robots)", r.CrawlAllowed)
		}
	})

	t.Run("done falls back to verdict_reason", func(t *testing.T) {
		// DONE has no SkipReason; reason must fall back to VerdictReason.
		a := &session.Article{URL: "u", Host: "h", State: session.DONE, VerdictReason: "hallucinated who"}
		r := Render(a, nil)
		if r == nil {
			t.Fatal("Render(DONE) = nil, want audit record")
		}
		if r.Status != "DONE" || r.Reason != "hallucinated who" {
			t.Errorf("DONE audit fields wrong: %+v", r)
		}
		if r.Event6 != nil || r.MediaName != "" {
			t.Errorf("DONE audit leaked collection fields: %+v", r)
		}
	})
}
