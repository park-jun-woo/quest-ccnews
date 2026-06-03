//ff:type feature=robots type=model
//ff:what 파싱된 robots.txt 룰셋. 선언된 순서대로의 user-agent 그룹 목록(순수 데이터).

package robots

// Ruleset is a parsed robots.txt: the ordered groups it declares.
type Ruleset struct {
	Groups []Group
}
