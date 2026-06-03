//ff:func feature=ingestion type=helper control=sequence
//ff:what SelectForward가 fetchPaths 에러를 그대로 전파하는지 검증한다.

package ingest

import (
	"errors"
	"testing"
)

func TestSelectForwardError(t *testing.T) {
	cur := Month{2026, 6}
	want := errors.New("boom")
	fetch := fakeFetch(nil, map[Month]error{cur: want})
	_, _, err := SelectForward(cur, map[string]bool{}, fetch)
	if !errors.Is(err, want) {
		t.Errorf("err = %v, want boom", err)
	}
}
