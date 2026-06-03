//ff:func feature=extract type=helper control=sequence level=error
//ff:what 기사 HTML 바이트를 받아 JSON-LD→OG 순으로 구조화 필드를 뽑고, 앵커 대상 본문 텍스트(articleBody 우선, 없으면 태그제거 전문)를 채워 Result로 돌려준다. 순수 함수, 네트워크 없음.

package extract

// Parse turns one article's raw WARC HTML into a Result. Source priority:
//
//  1. JSON-LD with an Article object  → Source "jsonld".
//  2. else OG/meta with an og:title   → Source "og".
//  3. else no structured data         → Source "" (gate will SKIP).
//
// The anchor-target BodyText is the JSON-LD articleBody when present, otherwise
// the tag-stripped full HTML — always computed so the gate can measure it even
// when no structured fields exist. Pure: HTML bytes in, Result out, no IO.
func Parse(htmlBytes []byte) Result {
	scripts := jsonLDScripts(htmlBytes)
	stripped := stripHTML(htmlBytes)

	var r Result
	if f, ok := extractJSONLD(scripts); ok {
		r.Fields = f
		r.Source = "jsonld"
	} else if f, ok := extractMeta(htmlBytes); ok {
		r.Fields = f
		r.Source = "og"
	}

	if r.Fields.Body != "" {
		r.BodyText = r.Fields.Body
	} else {
		r.BodyText = stripped
	}
	return r
}
