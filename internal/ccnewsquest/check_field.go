//ff:func feature=gate type=helper control=iteration dimension=1
//ff:what 한 present 필드의 앵커들을 원문에 대조한다(anchor.checkField의 textmatch 이식). textmatch.Normalize 후 빈/공백 앵커는 무효로 건너뛰고, 유효앵커가 0개면 statusUnanchored, 유효앵커 전부 textmatch.Contains면 statusAnchored, 하나라도 없으면 statusHallucination(앵커 표면형 반환). 순수 함수, anchor.Gate와 verdict 동치.

package ccnewsquest

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/textmatch"
)

// checkField mirrors anchor.checkField but matches via reins textmatch. Each anchor
// whose textmatch.Normalize form is empty (the "" / "  " cheese vector) is not a
// valid anchor: it is neither counted nor matched (Phase009 L0; textmatch.Contains
// already returns false for such tokens, but they must also be skipped from the
// valid count). The function returns the field status and, for a hallucination, the
// offending anchor's surface form (for the Fact). It is pure — unlike the original
// it does NOT mutate Field.Anchored, since the gate runs on a Submission and the
// ratchet persists state elsewhere.
func checkField(f *session.Field, source string) (fieldStatus, string) {
	valid := 0
	for _, a := range f.Anchors {
		if textmatch.Normalize(a) == "" {
			continue
		}
		valid++
		if !textmatch.Contains(source, a) {
			return statusHallucination, a
		}
	}
	if valid == 0 {
		return statusUnanchored, ""
	}
	return statusAnchored, ""
}
