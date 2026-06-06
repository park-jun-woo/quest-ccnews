//ff:func feature=robots type=helper control=sequence
//ff:what 동일 * 그룹 둘 중 한쪽만 Crawl-delay가 있을 때 병합 Group이 HasDelay·CrawlDelay를 이월(첫 명시 delay 우선)하는지 검증한다.

package robots

import "testing"

func TestSelectGroupMergeCarriesDelay(t *testing.T) {
	rs := Parse([]byte(
		"User-agent: *\n" +
			"Disallow:\n" +
			"\n" +
			"User-agent: *\n" +
			"Crawl-delay: 7\n" +
			"Disallow: /reporter/\n",
	))

	g := SelectGroup(rs, "parkjunwoo-quest")
	if g == nil {
		t.Fatal("expected merged group, got nil")
	}
	if !g.HasDelay {
		t.Fatalf("expected HasDelay true after merge, got %+v", g)
	}
	if g.CrawlDelay != 7 {
		t.Errorf("CrawlDelay = %d, want 7", g.CrawlDelay)
	}
}
