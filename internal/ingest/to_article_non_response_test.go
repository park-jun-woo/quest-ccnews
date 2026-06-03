//ff:func feature=ingestion type=helper control=sequence
//ff:what ToArticle가 response가 아닌 레코드에 대해 ok=false를 반환하는지 검증한다.

package ingest

import "testing"

func TestToArticleNonResponse(t *testing.T) {
	if _, ok := ToArticle("w", RecordView{Type: "request", TargetURI: "https://a.com/"}); ok {
		t.Error("non-response should yield ok=false")
	}
}
