//ff:func feature=output type=helper control=sequence
//ff:what BLOCKED/SKIPPED/DONE 기사를 audit 레코드(url/host/status/reason+crawl_allowed)로 만든다. reason은 skip_reason, 없으면(DONE) verdict_reason. 수집항목은 담지 않는다. 순수.

package output

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// renderAudit builds a BLOCKED/SKIPPED/DONE record with only the legality trail.
// reason comes from SkipReason (BLOCKED/SKIPPED); for DONE — which has no
// SkipReason — it falls back to the last failed attempt's VerdictReason.
func renderAudit(a *session.Article, host *session.Host) *Record {
	reason := a.SkipReason
	if reason == "" {
		reason = a.VerdictReason
	}
	r := &Record{
		URL:    a.URL,
		Host:   a.Host,
		Status: string(a.State),
		Reason: reason,
	}
	if host != nil && host.Robots != nil {
		r.CrawlAllowed = boolPtr(host.Robots.CrawlAllowed)
	}
	return r
}
