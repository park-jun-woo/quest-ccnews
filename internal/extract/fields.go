//ff:type feature=extract type=model
//ff:what 한 출처(JSON-LD 또는 OG/meta)에서 뽑은 구조화 필드 묶음. 빈 문자열은 "사이트가 선언하지 않음"을 뜻한다.

package extract

// Fields holds the structured fields pulled from one source (JSON-LD or
// OG/meta). Empty strings mean "not declared by the site".
type Fields struct {
	Title       string
	Author      string
	PublishedAt string
	MediaName   string
	Lang        string
	Body        string // JSON-LD articleBody; empty when not declared
}
