//ff:func feature=extract type=helper control=sequence level=error
//ff:what 추출+게이트를 기사 1건에 연동. PASS면 Extracted{Source,Title,Author,PublishedAt,BodyLen} 채우고 앵커용 본문 텍스트를 돌려준다. SKIP이면 State=SKIPPED+SkipReason 잠금. 상태 변경만, IO 없음.

package extract

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// Apply runs Parse + Gate on one article's WARC HTML and reflects the
// deterministic outcome onto the session Article. Only TODO articles are
// processed; an already-locked article is left untouched.
//
//   - PASS  → fills a.Extracted (Source, Title, Author, PublishedAt, BodyLen),
//     leaves State=TODO for the next gate (Phase006 anchoring), and returns the
//     anchor-target body text plus ok=true. The body text is NOT stored on the
//     session (only BodyLen is); the caller uses it for anchor verification.
//   - SKIPPED → locks State=SKIPPED with SkipReason, returns ok=false.
//
// State mutation only — no IO, no network.
func Apply(a *session.Article, htmlBytes []byte) (bodyText string, ok bool) {
	if a.State != session.TODO {
		return "", false
	}
	r := Parse(htmlBytes)
	d := Gate(r)
	if !d.Pass {
		a.State = session.SKIPPED
		a.SkipReason = d.SkipReason
		return "", false
	}
	a.Extracted = &session.Extracted{
		Title:       r.Fields.Title,
		Author:      r.Fields.Author,
		PublishedAt: r.Fields.PublishedAt,
		Source:      d.Source,
		BodyLen:     d.BodyLen,
	}
	if a.Lang == "" {
		a.Lang = r.Fields.Lang
	}
	return r.BodyText, true
}
