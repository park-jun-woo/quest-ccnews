//ff:func feature=output type=helper control=sequence
//ff:what session event6를 출력 event6로 매핑한다. 존재하는 필드만 renderField로 옮긴다. 순수.

package output

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// renderEvent6 maps the session event6 to the output event6, keeping only
// present fields and emitting anchors only when there are any.
func renderEvent6(ev *session.Event6) *Event6 {
	if ev == nil {
		return nil
	}
	return &Event6{
		Who:   renderField(ev.Who),
		When:  renderField(ev.When),
		Where: renderField(ev.Where),
		What:  renderField(ev.What),
		How:   renderField(ev.How),
		Why:   renderField(ev.Why),
	}
}
