//ff:func feature=output type=helper control=sequence
//ff:what renderCollected(PASS/REVIEW)가 전체 수집필드를 채우고, host/Extracted 유무에 따라 매체·라이선스·crawl_allowed·published_at을 선택적으로 처리하는지 검증한다.

package output

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestRenderCollected(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		a := &session.Article{
			URL:         "https://e.com/a",
			Host:        "e.com",
			State:       session.PASS,
			Lang:        "ko",
			CollectedAt: "2026-06-01T00:00:00Z",
			Extracted:   &session.Extracted{PublishedAt: "2026-05-31T12:00:00Z"},
			Event6: &session.Event6{
				Who:  &session.Field{Value: "Alice", Anchors: []string{"앨리스"}},
				What: &session.Field{Value: "won", Anchors: []string{"우승"}},
				Why:  &session.Field{Value: "talent"}, // present but no anchors
			},
		}
		host := &session.Host{
			MediaName: "Example News",
			SiteURL:   "https://e.com",
			License:   &session.License{Type: "CC-BY", URL: "https://cc/by", Source: "footer"},
			Robots:    &session.Robots{CrawlAllowed: true},
		}

		r := Render(a, host)
		if r == nil {
			t.Fatal("Render() = nil, want record")
		}
		if r.Status != "PASS" || r.URL != a.URL || r.Host != a.Host || r.Lang != "ko" {
			t.Errorf("basic fields wrong: %+v", r)
		}
		if r.CollectedAt != a.CollectedAt {
			t.Errorf("CollectedAt = %q", r.CollectedAt)
		}
		if r.PublishedAt != "2026-05-31T12:00:00Z" {
			t.Errorf("PublishedAt = %q (from Extracted)", r.PublishedAt)
		}
		if r.MediaName != "Example News" || r.SiteURL != "https://e.com" {
			t.Errorf("host fields wrong: %+v", r)
		}
		if r.License == nil || r.License.Type != "CC-BY" || r.License.URL != "https://cc/by" || r.License.Source != "footer" {
			t.Errorf("License = %+v", r.License)
		}
		if r.CrawlAllowed == nil || *r.CrawlAllowed != true {
			t.Errorf("CrawlAllowed = %v, want *true", r.CrawlAllowed)
		}
		if r.AnchorSummary != "2/3" {
			t.Errorf("AnchorSummary = %q, want 2/3", r.AnchorSummary)
		}
		if r.Event6 == nil || r.Event6.Who == nil || r.Event6.When != nil {
			t.Errorf("Event6 = %+v", r.Event6)
		}
	})

	t.Run("no host no extracted", func(t *testing.T) {
		a := &session.Article{URL: "u", Host: "h", State: session.REVIEW}
		r := Render(a, nil)
		if r == nil {
			t.Fatal("Render() = nil")
		}
		if r.Status != "REVIEW" {
			t.Errorf("Status = %q", r.Status)
		}
		if r.MediaName != "" || r.License != nil || r.CrawlAllowed != nil {
			t.Errorf("host fields should be empty: %+v", r)
		}
		if r.PublishedAt != "" {
			t.Errorf("PublishedAt = %q, want empty", r.PublishedAt)
		}
		if r.AnchorSummary != "0/0" || r.Event6 != nil {
			t.Errorf("empty event6: summary=%q event6=%+v", r.AnchorSummary, r.Event6)
		}
	})

	t.Run("host no license no robots", func(t *testing.T) {
		a := &session.Article{URL: "u", Host: "h", State: session.PASS}
		host := &session.Host{MediaName: "M"} // License nil, Robots nil
		r := Render(a, host)
		if r.MediaName != "M" {
			t.Errorf("MediaName = %q", r.MediaName)
		}
		if r.License != nil {
			t.Errorf("License = %+v, want nil", r.License)
		}
		if r.CrawlAllowed != nil {
			t.Errorf("CrawlAllowed = %v, want nil", r.CrawlAllowed)
		}
	})
}
