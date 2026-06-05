//ff:func feature=ingestion type=helper control=sequence
//ff:what blockArticle 단위테스트. guard가 nil이거나 호스트가 비면 Item을 무수정(TODO 유지)하고, guard가 허용하면 무수정, guard가 거부하면 payload State=BLOCKED·SkipReason 보존과 함께 Item을 quest.Apply(OutBlock)로 BLOCKED·now 스탬프 시드하는지 검증한다.

package runcmd

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/robots"
	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestBlockArticle(t *testing.T) {
	const now = "2026-06-05T00:00:00Z"

	t.Run("nil guard leaves the item untouched", func(t *testing.T) {
		a := &session.Article{URL: "https://ex.com/a", Host: "ex.com"}
		it := &quest.Item{Key: a.URL, State: quest.TODO}
		blockArticle(it, a, nil, now)
		if it.State != quest.TODO || a.State == session.BLOCKED {
			t.Errorf("State = %q / payload %q, want TODO untouched", it.State, a.State)
		}
	})

	t.Run("empty host skips the guard", func(t *testing.T) {
		a := &session.Article{URL: "https://deny.com/x", Host: ""}
		it := &quest.Item{Key: a.URL, State: quest.TODO}
		blockArticle(it, a, guardDenyingHost("deny.com"), now)
		if it.State != quest.TODO {
			t.Errorf("State = %q, want TODO (host empty → guard skipped)", it.State)
		}
	})

	t.Run("allowed host leaves the item TODO", func(t *testing.T) {
		a := &session.Article{URL: "https://ex.com/ok", Host: "ex.com"}
		it := &quest.Item{Key: a.URL, State: quest.TODO}
		// Empty ruleset → default-allow, purely from cache (no network).
		blockArticle(it, a, cachedGuard("ex.com", &robots.Ruleset{}), now)
		if it.State != quest.TODO || a.SkipReason != "" {
			t.Errorf("State = %q / reason %q, want TODO untouched", it.State, a.SkipReason)
		}
	})

	t.Run("denied host seeds the item BLOCKED with the reason stamped", func(t *testing.T) {
		a := &session.Article{URL: "https://deny.com/x", Host: "deny.com"}
		it := &quest.Item{Key: a.URL, State: quest.TODO}
		blockArticle(it, a, guardDenyingHost("deny.com"), now)
		if it.State != quest.BLOCKED {
			t.Errorf("State = %q, want BLOCKED", it.State)
		}
		if a.State != session.BLOCKED || a.SkipReason == "" {
			t.Errorf("payload State=%q reason=%q, want BLOCKED + reason", a.State, a.SkipReason)
		}
		if it.CollectedAt != now {
			t.Errorf("CollectedAt = %q, want %q (Apply stamp)", it.CollectedAt, now)
		}
	})
}
