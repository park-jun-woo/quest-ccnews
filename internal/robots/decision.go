//ff:type feature=robots type=model
//ff:what robots path 평가 결과. 허용 여부와 감사용 매칭 룰 문자열(예: "Disallow: /private")을 담는다.

package robots

// Decision is the outcome of evaluating a path against a ruleset: whether crawl
// is allowed and, for audit, the matched rule string ("" when no rule applied).
type Decision struct {
	Allowed bool
	Rule    string // e.g. "Disallow: /private/*" — empty when default-allow
}
