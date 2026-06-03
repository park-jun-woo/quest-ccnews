//ff:type feature=anchor type=model
//ff:what anchor 패키지 개요 — event6의 원어 앵커 토큰이 원문 본문에 substring으로 실재하는지(유효앵커 기준) 결정론적으로 대조해 PASS/REVIEW/FAIL을 낸다. value는 앵커 매핑 대상은 아니나 내재적 위생(플레이스홀더/길이)은 검사한다. 순수 로직.

// Package anchor is the deterministic fact-anchor gate (Phase006, hardened by
// Phase009). It checks that every anchor token an AI claimed for an event6 field
// appears verbatim (as a substring) in the article's original-language body text,
// blocking hallucinated facts. Generation is the AI's job; the verdict is the
// machine's.
//
// What the gate inspects:
//
//   - Field.Anchors (original-language surface forms) against the source body
//     text. Only "valid" anchors count: an anchor whose normalized form is empty
//     (the "" / whitespace cheese vector) is neither counted nor matched
//     (Phase009 L0), since strings.Contains(src,"") is trivially true.
//   - Field.Value's intrinsic hygiene (Phase009 L3): a present field's value must
//     not be empty, shorter than two runes, or a placeholder token
//     (Subject/Event/Unknown/N/A/…). This is a value-INTRINSIC check on the value
//     alone — it is NEVER compared to the anchors, so no anchor→value mapping is
//     attempted (the value's English/ISO rendering vs. original-language anchors
//     have no deterministic correspondence).
//   - Field.Interpretive is a display label only, never a gate input. REVIEW is
//     triggered solely by zero valid anchors on a present optional field.
//
// Normalization is applied identically to both the source text and each anchor
// (whitespace collapse) so matching is a symmetric surface-form comparison, not
// an inferred mapping.
//
// Verdict (machine-only, no opinion):
//
//   - PASS:   required who/when/what each have a hygienic value, ≥1 valid anchor,
//     and every valid anchor a source substring; present optional fields
//     (where/how/why) likewise have a hygienic value and every valid anchor a
//     source substring.
//   - REVIEW: required three PASS, but some present optional field has zero valid
//     anchors (a value with zero verifiable fact tokens — structurally
//     unverifiable, so flagged for human review).
//   - FAIL:   a required field is missing, any present field's value fails the
//     intrinsic hygiene check, or any present field has a valid anchor that is
//     absent from the source (hallucination). FAIL is not an article state — it
//     is a failed attempt the caller retries (Phase006 §제출 판정).
//
// The gate is a pure function: Event6 + body text in, a Result out, no IO. It
// also fills each present Field.Anchored with its per-field machine verdict.
package anchor
