//ff:type feature=article type=model
//ff:what 작업 목록의 기사 한 건(퀘스트). WARC 위치, 추출 결과, event6, 상태/판정/시도 로그를 담는다.

package session

// Article: one article quest in the work-list.
type Article struct {
	WARC          *WARCLoc   `json:"warc,omitempty"`
	URL           string     `json:"url"`            // WARC Target-URI
	Host          string     `json:"host"`           // host key into Session.Hosts
	Lang          string     `json:"lang,omitempty"` // source language code (output is English)
	State         State      `json:"state"`
	SkipReason    string     `json:"skip_reason,omitempty"` // recorded on BLOCKED/SKIPPED
	Tries         int        `json:"tries"`
	Extracted     *Extracted `json:"extracted,omitempty"`
	Event6        *Event6    `json:"event6,omitempty"`
	Verdict       string     `json:"verdict,omitempty"`
	VerdictReason string     `json:"verdict_reason,omitempty"`
	CollectedAt   string     `json:"collected_at,omitempty"`
	Log           []Attempt  `json:"log,omitempty"`

	// Emitted records that this article's terminal-state record has already been
	// appended to the JSONL output, so the sweep emits each article at most once
	// (Phase007). Additive field, default false — old sessions deserialize fine.
	Emitted bool `json:"emitted,omitempty"`
}
