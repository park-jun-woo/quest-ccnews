//ff:func feature=ingestion type=helper control=sequence
//ff:what ToArticle가 host 없는 Target-URI에 대해 ok=false를 반환하는지 검증한다.

package ingest

import "testing"

func TestToArticleNoHost(t *testing.T) {
	if _, ok := ToArticle("w", RecordView{Type: "response", TargetURI: "/no/host"}); ok {
		t.Error("hostless URI should yield ok=false")
	}
}
