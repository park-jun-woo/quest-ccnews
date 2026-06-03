//ff:func feature=robots type=helper control=selection
//ff:what 파싱된 한 줄(field,value)을 상태에 반영한다. user-agent/allow/disallow/crawl-delay별로 그룹·규칙·딜레이를 갱신한다.

package robots

// feed applies one parsed robots.txt line to the parser state, dispatching on
// the field name. Unknown fields end the current user-agent run. The detailed
// mutation for each field lives in single-purpose helpers to keep this switch
// flat.
func (p *parser) feed(field, value string) {
	switch field {
	case "user-agent":
		p.feedAgent(value)
	case "allow":
		p.feedRule(true, value)
	case "disallow":
		p.feedRule(false, value)
	case "crawl-delay":
		p.feedDelay(value)
	default:
		p.expectingAgent = false // unknown field ends the agent run
	}
}
