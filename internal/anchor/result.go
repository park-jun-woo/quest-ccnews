//ff:type feature=anchor type=model
//ff:what 앵커 게이트 1회 판정 결과. Verdict(PASS/REVIEW/FAIL)와 사람이 읽을 수 있는 Reason(사실 그대로). 호출측은 이걸로 상태 전이/재시도를 결정한다.

package anchor

// Result is one anchor-gate evaluation. Reason states the fact behind the
// verdict (which field, which anchor) — never an opinion. The gate fills each
// present Field.Anchored as a side effect so the caller can persist it.
type Result struct {
	Verdict Verdict
	Reason  string
}
