//ff:func feature=anchor type=helper control=iteration dimension=1
//ff:what 한 present 필드의 앵커들을 정규화된 원문에 대조한다. 앵커가 0개면 unanchored(미검증), 전부 substring이면 anchored, 하나라도 없으면 hallucination(앵커 텍스트 반환). 순수 함수, Field.Anchored를 채운다.

package anchor

import (
	"strings"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// checkField normalizes each anchor and tests it as a substring of the already
// normalized source. It returns the field status and, for a hallucination, the
// offending anchor's surface form (for the Reason). It sets f.Anchored true only
// when every anchor matched and there was at least one anchor. Pure: reads f,
// writes f.Anchored.
func checkField(f *session.Field, normSource string) (fieldStatus, string) {
	if len(f.Anchors) == 0 {
		f.Anchored = false
		return unanchored, ""
	}
	for _, a := range f.Anchors {
		if !strings.Contains(normSource, normalize(a)) {
			f.Anchored = false
			return hallucination, a
		}
	}
	f.Anchored = true
	return anchored, ""
}
