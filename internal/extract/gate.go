//ff:func feature=extract type=helper control=sequence
//ff:what 파싱된 Result를 신뢰 게이트로 판정한다. 구조화 출처/제목 없으면 SkipNoStructured, 본문<MinBodyLen이면 SkipBodyTooShort, 둘 다 충족하면 PASS. 순수 함수.

package extract

// Gate evaluates a parsed Result against the trust gate (Phase005):
//
//   - no structured source / no self-declared title → SKIPPED (SkipNoStructured)
//   - structured title present but BodyText < MinBodyLen → SKIPPED (SkipBodyTooShort)
//   - structured title present and BodyText >= MinBodyLen → PASS
//
// Pure — measures field presence and text length only, no opinion.
func Gate(r Result) Decision {
	d := Decision{Source: r.Source, BodyLen: len(r.BodyText)}
	if r.Source == "" || !r.Fields.hasArticle() {
		d.SkipReason = SkipNoStructured
		return d
	}
	if d.BodyLen < MinBodyLen {
		d.SkipReason = SkipBodyTooShort
		return d
	}
	d.Pass = true
	return d
}
