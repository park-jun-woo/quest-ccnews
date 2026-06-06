//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what 같은 종류의 그룹들(RFC 9309 §2.2.1 동일 에이전트 레코드)을 선언 순서대로 하나의 새 Group으로 병합한다. 순수 함수.

package robots

// mergeGroups combines several groups of the same kind (the same matching
// user-agent under RFC 9309 §2.2.1) into a single new Group, leaving the inputs
// untouched. Rules and Agents are concatenated in declaration order (union), so
// Evaluate's longest-match / Allow-first precedence runs over the merged rule
// set. The first explicitly-declared crawl-delay (HasDelay==true) wins and is
// carried over — fetchOK reads SelectGroup's HasDelay/CrawlDelay, so dropping it
// would silently lose the crawl-delay. Returns nil for an empty slice to keep
// the nil contract. Pure — no IO.
func mergeGroups(groups []*Group) *Group {
	if len(groups) == 0 {
		return nil
	}
	merged := &Group{}
	for _, g := range groups {
		merged.Agents = append(merged.Agents, g.Agents...)
		merged.Rules = append(merged.Rules, g.Rules...)
		if g.HasDelay && !merged.HasDelay {
			merged.HasDelay = true
			merged.CrawlDelay = g.CrawlDelay
		}
	}
	return merged
}
