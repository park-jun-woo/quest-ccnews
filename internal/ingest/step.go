//ff:type feature=ingestion type=model
//ff:what 한 번의 인제스천 스텝 결과. 처리한 WARC 이름과 추가된 기사 수, 그리고 트랙이 멈춰야 하는지(waiting/exhausted)와 그 상태.

package ingest

// StepResult is the outcome of advancing one track by a single WARC.
type StepResult struct {
	WarcName    string // the WARC just processed (empty when no work was done)
	ArticlesAdd int    // number of TODO articles appended from that WARC
	Stopped     bool   // true when the track reached a terminal-for-now state
	State       string // when Stopped: StateWaiting (forward) or StateExhausted (backward)
}
