//ff:func feature=robots type=helper control=sequence
//ff:what URL 또는 경로에서 평가용 path(+query)를 뽑아 정규화한다. 빈 값은 "/"로. 순수 함수.

package robots

import "strings"

// NormalizePath extracts the path (with query, per RFC 9309 §2.2.2 the rule is
// matched against path plus query) from a raw URL or bare path and normalizes
// it for comparison. A full URL has its scheme/host stripped; a bare path is
// used as-is. An empty result becomes "/". Pure — no IO.
func NormalizePath(rawURLOrPath string) string {
	p := rawURLOrPath
	if strings.Contains(rawURLOrPath, "://") {
		p = pathWithQuery(rawURLOrPath)
	}
	if p == "" || !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	return p
}
