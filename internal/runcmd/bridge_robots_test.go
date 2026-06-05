//ff:func feature=ingestion type=helper control=sequence
//ff:what bridge의 robots 경로(blockArticle) 단위테스트. robots 거부 기사는 BLOCKED로 직접 시드되며 payload State=BLOCKED·SkipReason·Item CollectedAt(Apply 스탬프)이 보존되는지, 호스트가 빈 기사는 guard를 건너뛰어 TODO로 남는지 검증한다.

package runcmd

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestBridgeRobots(t *testing.T) {
	const now = "2026-06-05T00:00:00Z"

	t.Run("robots-denied article is seeded straight to BLOCKED", func(t *testing.T) {
		scratch := session.New("ua", "cc-news")
		a := &session.Article{URL: "https://deny.com/x", Host: "deny.com"}
		scratch.Articles = []*session.Article{a}
		s := quest.New()
		guard := guardDenyingHost("deny.com")

		n := bridge(scratch, s, guard, now)

		if n != 1 {
			t.Fatalf("seeded = %d, want 1", n)
		}
		it := s.Items[0]
		if it.State != quest.BLOCKED {
			t.Errorf("State = %q, want BLOCKED", it.State)
		}
		if a.State != session.BLOCKED {
			t.Errorf("payload State = %q, want BLOCKED", a.State)
		}
		if a.SkipReason == "" {
			t.Errorf("payload SkipReason empty, want robots reason")
		}
		if it.CollectedAt != now {
			t.Errorf("CollectedAt = %q, want %q (Apply stamp)", it.CollectedAt, now)
		}
	})

	t.Run("guard is consulted only when article has a host", func(t *testing.T) {
		scratch := session.New("ua", "cc-news")
		// Denying guard, but article has no Host → guard skipped, seeded TODO.
		scratch.Articles = []*session.Article{{URL: "https://deny.com/y", Host: ""}}
		s := quest.New()
		guard := guardDenyingHost("deny.com")

		bridge(scratch, s, guard, now)

		if s.Items[0].State != quest.TODO {
			t.Errorf("State = %q, want TODO (host empty → guard skipped)", s.Items[0].State)
		}
	})
}
