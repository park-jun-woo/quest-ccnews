//ff:func feature=ingestion type=command control=sequence level=error
//ff:what scratch를 reins 세션으로 브리지하고 세션을 저장한다. ingest의 매-WARC Save 콜백(중단 시 커서 재개)과 루프 종료 후 1회에서 재사용된다. bridge로 새 Item을 시드한 뒤 세션을 sessionPath에 저장하고, 새로 시드된 수>0이면 seed 보고 한 줄을 w에 출력한다.

package runcmd

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// flush bridges the scratch into the reins session and persists it. Reused by the
// ingest Save callback (per WARC) and once more after the loop returns.
func (o *options) flush(scratch *session.Session, s *quest.Session, guard *robotsGuard, now string, w io.Writer) error {
	n := bridge(scratch, s, guard, now)
	if err := s.Save(o.sessionPath()); err != nil {
		return err
	}
	if n > 0 {
		fmt.Fprintf(w, "seed: +%d items → %s\n", n, o.sessionPath())
	}
	return nil
}
