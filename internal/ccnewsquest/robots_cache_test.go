//ff:func feature=robots type=helper control=sequence
//ff:what newRobotsCache 단위 테스트. 빈 pick-time robots 캐시를 생성하는지 검증한다. rulesets 맵이 nil이 아니고(즉시 적재 가능) 비어 있어야 하며, 매 호출이 독립 인스턴스를 돌려줘야 한다(프로세스별 1회 생성·공유는 Def 책임).
package ccnewsquest

import "testing"

func TestNewRobotsCache(t *testing.T) {
	c := newRobotsCache()
	if c == nil {
		t.Fatal("newRobotsCache returned nil")
	}
	if c.rulesets == nil {
		t.Fatal("rulesets map is nil; allowed() would panic on its first cache write")
	}
	if len(c.rulesets) != 0 {
		t.Fatalf("fresh cache has %d rulesets, want 0", len(c.rulesets))
	}

	// The map is writable (the zero-value would not be).
	c.rulesets["h.com"] = nil
	if len(c.rulesets) != 1 {
		t.Fatalf("ruleset write did not take: len=%d", len(c.rulesets))
	}

	// Each call is a fresh, independent instance.
	other := newRobotsCache()
	if len(other.rulesets) != 0 {
		t.Fatalf("second cache is not independent: len=%d", len(other.rulesets))
	}
}
