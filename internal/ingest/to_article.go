//ff:func feature=ingestion type=helper control=sequence
//ff:what WARC response 레코드 뷰를 TODO 상태의 session.Article로 변환한다. response가 아니거나 host 없으면 ok=false. 순수 함수.

package ingest

import (
	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

// ToArticle converts a WARC record view into a fresh TODO article located by
// {warcFile, offset}. It returns ok=false for non-response records or records
// whose Target-URI has no usable host, so the caller can skip them. Pure (no IO):
// the body text is never stored — the WARC at warc{file,offset} remains the
// source of truth. (Phase002 design decision 1.)
func ToArticle(warcFile string, rv RecordView) (*session.Article, bool) {
	if rv.Type != "response" {
		return nil, false
	}
	host, ok := HostOf(rv.TargetURI)
	if !ok {
		return nil, false
	}
	return &session.Article{
		WARC:  &session.WARCLoc{File: warcFile, Offset: rv.Offset},
		URL:   rv.TargetURI,
		Host:  host,
		State: session.TODO,
	}, true
}
