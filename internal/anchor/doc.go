//ff:type feature=anchor type=model
//ff:what anchor 패키지 개요 — event6의 원어 앵커 토큰이 원문 본문에 substring으로 실재하는지 결정론적으로 대조해 PASS/REVIEW/FAIL을 낸다. value는 게이트 대상 아님. 순수 로직.

// Package anchor is the deterministic fact-anchor gate (Phase006). It checks
// that every anchor token an AI claimed for an event6 field appears verbatim
// (as a substring) in the article's original-language body text, blocking
// hallucinated facts. Generation is the AI's job; the verdict is the machine's.
//
// What the gate inspects:
//
//   - Only Field.Anchors (original-language surface forms) against the source
//     body text. Field.Value (English/ISO output) is a display rendering and is
//     NEVER a gate target — no anchor→value mapping is attempted.
//   - Field.Interpretive is a display label only, never a gate input. REVIEW is
//     triggered solely by len(Anchors)==0 on a present optional field.
//
// Normalization is applied identically to both the source text and each anchor
// (whitespace collapse) so matching is a symmetric surface-form comparison, not
// an inferred mapping.
//
// Verdict (machine-only, no opinion):
//
//   - PASS:   required who/when/what each have a value, a non-empty Anchors, and
//     every anchor is a source substring; present optional fields (where/how/why)
//     likewise have every anchor as a source substring.
//   - REVIEW: required three PASS, but some present optional field has
//     len(Anchors)==0 (a value with zero verifiable fact tokens — structurally
//     unverifiable, so flagged for human review).
//   - FAIL:   a required field is missing, or any present field has an anchor
//     that is absent from the source (hallucination). FAIL is not an article
//     state — it is a failed attempt the caller retries (Phase006 §제출 판정).
//
// The gate is a pure function: Event6 + body text in, a Result out, no IO. It
// also fills each present Field.Anchored with its per-field machine verdict.
package anchor
