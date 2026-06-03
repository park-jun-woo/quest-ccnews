//ff:type feature=ingestion type=model
//ff:what 인제스천 run 루프 설정. 처리할 트랙 목록, 최대 처리 WARC 수(0=무제한), 클럭, 세션 저장 경로.

package ingest

import "time"

// RunOptions configures the ingestion run loop.
type RunOptions struct {
	Tracks   []Track   // forward / backward / both
	MaxWarcs int       // stop after this many WARCs processed (0 = unbounded)
	Now      time.Time // clock for the newest-month default (zero → time.Now)
	Save     func() error
}
