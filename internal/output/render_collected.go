//ff:func feature=output type=helper control=sequence
//ff:what PASS/REVIEW 기사를 전체 수집항목 레코드로 만든다. host/Extracted가 있으면 매체·라이선스·crawl_allowed·published_at을 채운다. 순수.

package output

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// renderCollected builds a PASS/REVIEW record with the full collection fields.
func renderCollected(a *session.Article, host *session.Host) *Record {
	r := &Record{
		URL:           a.URL,
		Host:          a.Host,
		Status:        string(a.State),
		Lang:          a.Lang,
		CollectedAt:   a.CollectedAt,
		Event6:        renderEvent6(a.Event6),
		AnchorSummary: anchorSummary(a.Event6),
	}
	if a.Extracted != nil {
		r.PublishedAt = a.Extracted.PublishedAt
	}
	if host != nil {
		r.MediaName = host.MediaName
		r.SiteURL = host.SiteURL
		if host.License != nil {
			r.License = &FieldLicense{
				Type:   host.License.Type,
				URL:    host.License.URL,
				Source: host.License.Source,
			}
		}
		if host.Robots != nil {
			r.CrawlAllowed = boolPtr(host.Robots.CrawlAllowed)
		}
	}
	return r
}
