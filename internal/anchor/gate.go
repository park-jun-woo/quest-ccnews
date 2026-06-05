//ff:func feature=anchor type=helper control=iteration dimension=1 level=error
//ff:what event6 1회 제출을 원문 본문에 대조해 PASS/REVIEW/FAIL을 판정하는 순수 게이트. 필수(who/what)는 위생적 value+유효앵커 전부 substring, 선택(when/where/how/why)은 present 시 위생적 value+유효앵커 substring. 플레이스홀더 value 또는 환각=FAIL, present 선택필드 유효앵커0개=REVIEW. Field.Anchored를 채운다.

package anchor

import (
	"fmt"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// Gate evaluates one event6 submission against the source body text and returns
// the deterministic verdict (Phase006 §제출 판정). It is a pure function over
// (ev, source): the only mutation is filling each present Field.Anchored with
// its per-field machine verdict. No IO.
//
// Required who/what must each be present with a hygienic Value (Phase009 L3:
// non-empty, ≥2 runes, not a placeholder), ≥1 valid anchor, and every valid anchor
// a source substring. A missing required field (nil), a placeholder/empty Value,
// zero valid anchors, or any hallucinated anchor in any field is FAIL. A present
// optional field (when/where/how/why) with zero valid anchors is structurally
// unverifiable → REVIEW (only when the required two otherwise PASS). A "valid"
// anchor is one whose normalized form is non-empty (Phase009 L0).
//
// Per-field checking is delegated to checkRequired/checkOptional so this gate
// stays a flat traversal of the two field lists.
func Gate(ev *session.Event6, source string) Result {
	if ev == nil {
		return Result{Verdict: FAIL, Reason: "event6 없음"}
	}
	normSource := normalize(source)

	required := []namedField{
		{"who", ev.Who}, {"what", ev.What},
	}
	optional := []namedField{
		{"when", ev.When}, {"where", ev.Where}, {"how", ev.How}, {"why", ev.Why},
	}

	// Required fields: a placeholder/empty value or zero valid anchors is FAIL
	// ("부재" 변명 불가).
	for _, nf := range required {
		if res := checkRequired(nf, normSource); res != nil {
			return *res
		}
	}

	// Optional fields: present (non-nil) ones are checked. A placeholder value or
	// a hallucinated anchor is FAIL; a present field with zero valid anchors defers
	// the verdict to REVIEW.
	review := false
	var reviewName string
	for _, nf := range optional {
		res, unanchored := checkOptional(nf, normSource)
		if res != nil {
			return *res
		}
		if unanchored && !review {
			review = true
			reviewName = nf.name
		}
	}

	if review {
		return Result{Verdict: REVIEW, Reason: fmt.Sprintf("선택 필드 %s 유효앵커 0개 — 구조적 미검증(사람 확인 필요)", reviewName)}
	}
	return Result{Verdict: PASS, Reason: "필수·선택 앵커 전부 원문 substring 확인"}
}
