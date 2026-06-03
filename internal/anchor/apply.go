//ff:func feature=anchor type=helper control=selection
//ff:what 게이트 판정을 기사 상태기계에 적용한다. PASS/REVIEW면 event6 잠금+Verdict/Reason/CollectedAt 기록, FAIL이면 tries++ 후 TODO 유지(MaxTries 초과 시 DONE). Attempt 로그를 항상 남긴다. 순수 상태 변이(now는 주입).

package anchor

import "github.com/park-jun-woo/quest-ccnews/internal/session"

// Apply transitions an article per a gate Result (Phase006 §제출 판정). It is a
// pure state mutation — no IO; collectedAt (the lock timestamp) is injected so it
// stays testable. PASS/REVIEW lock the article (state + verdict + event6 stays as
// submitted, with Field.Anchored already filled by Gate) and stamp CollectedAt.
// FAIL is a failed attempt: tries++ and the article stays TODO for retry, locking
// to DONE once tries reaches session.MaxTries (the constant's documented "after
// this many failed attempts"). Every call appends an Attempt log entry with the
// verdict and its reason.
//
// The submitted event6 is attached to the article before applying so PASS/REVIEW
// persist exactly what the gate evaluated.
func Apply(a *session.Article, ev *session.Event6, res Result, collectedAt string) {
	a.Event6 = ev
	a.Log = append(a.Log, session.Attempt{
		Try:     a.Tries + 1,
		Verdict: string(res.Verdict),
		Reason:  res.Reason,
	})

	switch res.Verdict {
	case PASS, REVIEW:
		if res.Verdict == PASS {
			a.State = session.PASS
		} else {
			a.State = session.REVIEW
		}
		a.Verdict = string(res.Verdict)
		a.VerdictReason = res.Reason
		a.CollectedAt = collectedAt
	default: // FAIL
		a.Tries++
		if a.Tries >= session.MaxTries {
			a.State = session.DONE
			a.Verdict = string(res.Verdict)
			a.VerdictReason = res.Reason
		}
	}
}
