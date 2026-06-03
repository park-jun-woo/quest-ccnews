//ff:func feature=robots type=helper control=sequence
//ff:what 전체 URL에서 escaped path와 raw query를 합쳐 돌려준다. 파싱 실패 시 원본을 그대로 반환한다. 순수 함수.

package robots

import "net/url"

// pathWithQuery parses a full URL and returns its escaped path with the raw
// query appended ("?..."). On a parse error it falls back to the original
// string so the caller can still normalize it. Pure — no IO.
func pathWithQuery(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	p := u.EscapedPath()
	if u.RawQuery != "" {
		p += "?" + u.RawQuery
	}
	return p
}
