//ff:type feature=extract type=model
//ff:what 한 기사 HTML 파싱의 전체 결과. 선택된 구조화 필드 + 출처 라벨(jsonld|og|"") + 앵커 대상 본문 텍스트. 게이트 판정의 입력.

package extract

// Result is the full outcome of parsing one article's HTML: the chosen
// structured fields, the source label ("jsonld"|"og"|""), and the anchor-target
// body text (articleBody when present, else tag-stripped full HTML). The body
// text is returned for Phase006 anchoring and is never persisted in the session.
type Result struct {
	Fields Fields
	Source string // "jsonld" | "og" | "" (no structured data)
	// BodyText is the anchor-target text: Fields.Body when non-empty, else the
	// tag-stripped full HTML. Empty Source leaves this as the stripped HTML too,
	// so the gate can still report which SKIP reason applies.
	BodyText string
}
