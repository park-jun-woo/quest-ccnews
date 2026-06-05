//ff:func feature=robots type=helper control=sequence
//ff:what newRobotsGuard 단위테스트. client 생성, userAgent에서 추출한 productToken, hosts 맵이 참조 공유(보존 robots 캐시)되는지, rulesets가 초기화되는지 검증한다.

package runcmd

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestNewRobotsGuard(t *testing.T) {
	hosts := map[string]*session.Host{"ex.com": {MediaName: "Ex"}}
	g := newRobotsGuard("parkjunwoo-quest/0.1 (+url)", hosts)

	if g.client == nil {
		t.Errorf("client nil")
	}
	if g.productToken != "parkjunwoo-quest" {
		t.Errorf("productToken = %q, want parkjunwoo-quest", g.productToken)
	}
	// hosts is shared by reference (it is the persisted robots cache).
	if g.hosts == nil || g.hosts["ex.com"] == nil {
		t.Fatalf("hosts not shared: %v", g.hosts)
	}
	g.hosts["new.com"] = &session.Host{}
	if hosts["new.com"] == nil {
		t.Errorf("guard.hosts is not the same map as the passed-in cache")
	}
	if g.rulesets == nil {
		t.Errorf("rulesets not initialized")
	}
}
