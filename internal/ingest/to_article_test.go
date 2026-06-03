//ff:func feature=ingestion type=helper control=sequence
//ff:what ToArticle가 response 레코드를 소문자 host의 TODO Article로 변환하는지 검증한다.

package ingest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
)

func TestToArticleResponse(t *testing.T) {
	rv := RecordView{Type: "response", TargetURI: "https://News.Example.com/x", Offset: 7}
	a, ok := ToArticle("CC-NEWS-x.warc.gz", rv)
	if !ok {
		t.Fatal("want ok=true for response record")
	}
	if a.WARC == nil || a.WARC.File != "CC-NEWS-x.warc.gz" || a.WARC.Offset != 7 {
		t.Errorf("WARC loc = %+v", a.WARC)
	}
	if a.URL != "https://News.Example.com/x" {
		t.Errorf("URL = %q", a.URL)
	}
	if a.Host != "news.example.com" {
		t.Errorf("Host = %q, want lowercase", a.Host)
	}
	if a.State != session.TODO {
		t.Errorf("State = %q, want TODO", a.State)
	}
}
