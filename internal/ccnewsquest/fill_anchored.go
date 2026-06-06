//ff:func feature=gate type=helper control=iteration dimension=1
//ff:what 제출 event6의 각 present 필드에 게이트와 동일한 순수 함수(checkField/fieldStatus)로 anchored를 채운다(Phase013 D). event6Of와 동일한 필수(who/what)·선택(when/where/how/why) 필드 순서로 돌며, present 필드는 f.Anchored = (checkField(f, source)==statusAnchored)로 설정한다. verdict 계산에 쓰는 바로 그 함수·Source·필드순서를 재사용하므로 산출물의 anchored가 verdict와 항상 일치(동어반복적 동치, 별도 휴리스틱 아님). SetPayload 직전에 호출한다.
package ccnewsquest

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// fillAnchored sets ev.<field>.Anchored on every present field using the exact same
// pure function the gate uses (checkField/fieldStatus) over the same Source, so the
// persisted per-field anchored flag matches the verdict field-for-field (Phase013 D).
// checkField is intentionally non-mutating, so the gate left Field.Anchored untouched;
// this is the single place the verified flag is stamped before SetPayload. It walks
// who/what then when/where/how/why — the order event6Of/the rules use — and only
// touches present (non-nil) fields, leaving absent fields nil.
func fillAnchored(ev *session.Event6, source string) {
	for _, f := range []*session.Field{ev.Who, ev.What, ev.When, ev.Where, ev.How, ev.Why} {
		if f == nil {
			continue
		}
		status, _ := checkField(f, source)
		f.Anchored = status == statusAnchored
	}
}
