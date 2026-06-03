//ff:func feature=anchor type=helper control=iteration dimension=1
//ff:what 한 present 필드의 앵커들을 정규화된 원문에 대조한다. normalize 후 빈/공백인 앵커는 무효로 건너뛰고, 유효앵커가 0개면 unanchored(미검증), 유효앵커 전부 substring이면 anchored, 하나라도 없으면 hallucination(앵커 텍스트 반환). 순수 함수, Field.Anchored를 채운다.

package anchor

import (
	"strings"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// checkField normalizes each anchor and tests it as a substring of the already
// normalized source. An anchor whose normalized form is empty (the "" / "  "
// cheese vector — strings.Contains(src,"")==true) is not a valid anchor: it is
// neither counted nor matched (Phase009 L0). It returns the field status and,
// for a hallucination, the offending anchor's surface form (for the Reason). It
// sets f.Anchored true only when every valid anchor matched and there was at
// least one valid anchor. Pure: reads f, writes f.Anchored.
func checkField(f *session.Field, normSource string) (fieldStatus, string) {
	valid := 0
	for _, a := range f.Anchors {
		na := normalize(a)
		if na == "" {
			continue
		}
		valid++
		if !strings.Contains(normSource, na) {
			f.Anchored = false
			return hallucination, a
		}
	}
	if valid == 0 {
		f.Anchored = false
		return unanchored, ""
	}
	f.Anchored = true
	return anchored, ""
}
