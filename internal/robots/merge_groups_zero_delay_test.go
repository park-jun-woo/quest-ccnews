//ff:func feature=robots type=helper control=sequence
//ff:what mergeGroups가 명시적 crawl-delay 0(HasDelay true·CrawlDelay 0)을 효력 있는 delay로 이월하고 이후 비-0 delay가 덮어쓰지 못하게 하는지 검증한다.

package robots

import "testing"

func TestMergeGroupsZeroDelayIsExplicit(t *testing.T) {
	// HasDelay true with CrawlDelay 0 ("crawl-delay: 0") must carry, and a later
	// non-zero delay must not override it.
	a := &Group{Agents: []string{"*"}, HasDelay: true, CrawlDelay: 0}
	b := &Group{Agents: []string{"*"}, HasDelay: true, CrawlDelay: 4}
	g := mergeGroups([]*Group{a, b})
	if !g.HasDelay {
		t.Fatal("HasDelay = false, want true for explicit crawl-delay: 0")
	}
	if g.CrawlDelay != 0 {
		t.Errorf("CrawlDelay = %d, want 0 (first explicit delay wins)", g.CrawlDelay)
	}
}
