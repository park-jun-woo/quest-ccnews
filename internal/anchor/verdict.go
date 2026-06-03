//ff:type feature=anchor type=model
//ff:what 앵커 게이트가 event6 1회 제출을 원문에 대조한 결과 종류. PASS/REVIEW는 기사 상태로 잠그고, FAIL은 상태가 아니라 재시도 가능한 실패 시도다.

package anchor

// Verdict is the outcome of evaluating one event6 submission against the source.
type Verdict string

const (
	// PASS: required fields anchored and every present field's anchors are
	// source substrings → caller locks the article to state PASS.
	PASS Verdict = "PASS"
	// REVIEW: required fields PASS but a present optional field has no anchors
	// (structurally unverifiable) → caller locks the article to state REVIEW.
	REVIEW Verdict = "REVIEW"
	// FAIL: a required field is missing, or some anchor is absent from the
	// source (hallucination). Not an article state — a failed attempt the caller
	// retries (tries++; locks to DONE only when MaxTries is exhausted).
	FAIL Verdict = "FAIL"
)
