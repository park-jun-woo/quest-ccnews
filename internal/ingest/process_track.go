//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what 한 트랙을 WARC 한 개만큼 전진시킨다: 다음 WARC 선택→다운로드→스캔(기사 추가)→커서/processed 갱신. 더 받을 게 없으면 waiting(forward)/exhausted(backward)로 멈춤.

package ingest

import (
	"time"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// ProcessTrack advances one ingestion track by a single WARC. It selects the next
// unprocessed WARC (forward: newest month only; backward: walking toward the
// past), downloads it, scans its response records into TODO articles appended to
// the session, then ratchets the cursor and processed_warcs. When there is no
// WARC left to take it stops the track: StateWaiting for forward (poll later),
// StateExhausted for backward (reached 2016-08). The session is mutated in place;
// persistence (Save) is the caller's job. now drives the newest-month default for
// a fresh cursor.
func (c *Client) ProcessTrack(s *session.Session, track Track, now time.Time) (StepResult, error) {
	cur := EnsureCursor(&s.Ingestion, track)
	processed := ProcessedSet(s.Ingestion.ProcessedWarcs)
	start := StartMonth(cur, now)

	var (
		objectPath string
		done       bool
		err        error
	)
	if track == Backward {
		objectPath, done, err = SelectNext(start, processed, c.FetchPaths)
	} else {
		objectPath, done, err = SelectForward(start, processed, c.FetchPaths)
	}
	if err != nil {
		return StepResult{}, err
	}
	if done {
		state := StateWaiting
		if track == Backward {
			state = StateExhausted
		}
		cur.State = state
		return StepResult{Stopped: true, State: state}, nil
	}

	warcName := WarcName(objectPath)
	localPath, err := c.DownloadWarc(objectPath)
	if err != nil {
		return StepResult{}, err
	}

	added := 0
	scanErr := c.ScanWarc(localPath, warcName, func(a *session.Article) {
		s.Articles = append(s.Articles, a)
		added++
	})
	if scanErr != nil {
		return StepResult{}, scanErr
	}

	AdvanceCursor(&s.Ingestion, cur, warcName)
	return StepResult{WarcName: warcName, ArticlesAdd: added}, nil
}
