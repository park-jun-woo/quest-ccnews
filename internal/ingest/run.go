//ff:func feature=ingestion type=helper control=iteration dimension=2 level=error
//ff:what 인제스천 루프: 선택된 트랙을 라운드로빈으로 돌며 WARC를 한 개씩 처리하고 매번 세션을 저장한다. 모든 트랙이 멈추거나 MaxWarcs에 도달하면 끝낸다.

package ingest

import (
	"fmt"
	"io"
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// Run drives the two-track ingestion loop. It round-robins the requested tracks,
// advancing each by one WARC per turn (ProcessTrack), and persists the session
// after every step so an interrupt is always resumable from the cursor (ratchet).
// A forward track that reaches StateWaiting and a backward track at
// StateExhausted are dropped from the rotation; the loop ends when no track
// remains or MaxWarcs WARCs have been processed. Progress is written to w.
func Run(c *Client, s *session.Session, opt RunOptions, w io.Writer) error {
	now := opt.Now
	if now.IsZero() {
		now = time.Now()
	}
	active := append([]Track(nil), opt.Tracks...)
	processedCount := 0

	for len(active) > 0 {
		next := active[:0]
		for _, track := range active {
			if opt.MaxWarcs > 0 && processedCount >= opt.MaxWarcs {
				return nil
			}
			res, err := c.ProcessTrack(s, track, now)
			if err != nil {
				return fmt.Errorf("track %s: %w", track, err)
			}
			if err := saveSession(opt); err != nil {
				return err
			}
			if res.Stopped {
				fmt.Fprintf(w, "[%s] stopped: %s\n", track, res.State)
				continue // do not re-queue this track
			}
			fmt.Fprintf(w, "[%s] %s: +%d articles\n", track, res.WarcName, res.ArticlesAdd)
			processedCount++
			next = append(next, track)
		}
		active = next
	}
	return nil
}
