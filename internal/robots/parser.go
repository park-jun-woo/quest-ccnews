//ff:type feature=robots type=model
//ff:what robots.txt 파싱 누적 상태. 만들고 있는 룰셋, 현재 그룹 포인터, 직전 줄이 user-agent였는지 플래그를 담는다.

package robots

// parser accumulates Parse state across lines: the ruleset under construction,
// the current group pointer, and whether the previous non-blank line was a
// user-agent line (consecutive agents accumulate into one group).
type parser struct {
	rs             *Ruleset
	cur            *Group
	expectingAgent bool
}
