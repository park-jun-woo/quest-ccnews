//ff:type feature=output type=model
//ff:what JSONL 출력 레코드 스키마. collected(PASS/REVIEW)는 전체 수집항목+event6, audit(BLOCKED/SKIPPED)는 url/host/status/reason/crawl_allowed만.

package output

// Record is one JSONL line. collected (PASS/REVIEW) carries the full user-
// specified collection fields plus the event6 evidence; audit (BLOCKED/SKIPPED)
// carries only the legality trail. Body text is never emitted (facts only).
type Record struct {
	URL    string `json:"url"`
	Host   string `json:"host"`
	Status string `json:"status"` // PASS | REVIEW | BLOCKED | SKIPPED | DONE

	// collected (PASS/REVIEW) fields.
	MediaName     string        `json:"media_name,omitempty"`
	SiteURL       string        `json:"site_url,omitempty"`
	Lang          string        `json:"lang,omitempty"`
	PublishedAt   string        `json:"published_at,omitempty"`
	CollectedAt   string        `json:"collected_at,omitempty"`
	License       *FieldLicense `json:"license,omitempty"`
	AnchorSummary string        `json:"anchor_summary,omitempty"` // e.g. "4/4" (anchored / present)
	Event6        *Event6       `json:"event6,omitempty"`

	// audit (BLOCKED/SKIPPED) fields.
	Reason string `json:"reason,omitempty"` // skip_reason

	// CrawlAllowed appears in both shapes (user-specified field). Pointer so the
	// flag is always present even when false, without colliding with omitempty.
	CrawlAllowed *bool `json:"crawl_allowed,omitempty"`
}
