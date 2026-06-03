//ff:func feature=anchor type=helper control=iteration dimension=1 level=error
//ff:what event6 1회 제출을 원문 본문에 대조해 PASS/REVIEW/FAIL을 판정하는 순수 게이트. 필수(who/when/what)는 값+앵커 전부 substring, 선택(where/how/why)은 present 시 앵커 substring. 환각=FAIL, present 선택필드 앵커0개=REVIEW. Field.Anchored를 채운다.

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
// Required who/when/what must each be present with a non-empty Value, a non-empty
// Anchors, and every anchor a source substring. A missing required field (nil,
// empty Value, or empty Anchors) or any hallucinated anchor in any field is FAIL.
// A present optional field (where/how/why) with len(Anchors)==0 is structurally
// unverifiable → REVIEW (only when the required three otherwise PASS).
//
// Per-field checking is delegated to checkRequired/checkOptional so this gate
// stays a flat traversal of the two field lists.
func Gate(ev *session.Event6, source string) Result {
	if ev == nil {
		return Result{Verdict: FAIL, Reason: "event6 없음"}
	}
	normSource := normalize(source)

	required := []namedField{
		{"who", ev.Who}, {"when", ev.When}, {"what", ev.What},
	}
	optional := []namedField{
		{"where", ev.Where}, {"how", ev.How}, {"why", ev.Why},
	}

	// Required fields: missing value or missing anchors is FAIL ("부재" 변명 불가).
	for _, nf := range required {
		if res := checkRequired(nf, normSource); res != nil {
			return *res
		}
	}

	// Optional fields: present (non-nil) ones are checked. A hallucinated anchor
	// is FAIL; an anchorless present field defers the verdict to REVIEW.
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
		return Result{Verdict: REVIEW, Reason: fmt.Sprintf("선택 필드 %s 앵커 0개 — 구조적 미검증(사람 확인 필요)", reviewName)}
	}
	return Result{Verdict: PASS, Reason: "필수·선택 앵커 전부 원문 substring 확인"}
}
