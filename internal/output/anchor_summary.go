//ff:func feature=output type=helper control=iteration dimension=1
//ff:what 존재하는 event6 필드 중 앵커를 가진 필드 수/존재 필드 수를 "anchored/present"(예: "4/4")로 센다. 없으면 "0/0". 순수.

package output

import (
	"strconv"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// anchorSummary reports "anchored/present" — the count of present event6 fields
// that carry at least one anchor over the count of present fields (e.g. "4/4").
// Returns "0/0" when no fields are present.
func anchorSummary(ev *session.Event6) string {
	if ev == nil {
		return "0/0"
	}
	present, anchored := 0, 0
	for _, f := range []*session.Field{ev.Who, ev.When, ev.Where, ev.What, ev.How, ev.Why} {
		if f == nil {
			continue
		}
		present++
		if len(f.Anchors) > 0 {
			anchored++
		}
	}
	return strconv.Itoa(anchored) + "/" + strconv.Itoa(present)
}
