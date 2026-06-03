//ff:func feature=ingestion type=helper control=sequence
//ff:what SelectNext가 fetchPaths 에러를 그대로 전파하는지 검증한다.

package ingest

import (
	"errors"
	"testing"
)

func TestSelectNextError(t *testing.T) {
	start := Month{2026, 6}
	want := errors.New("net down")
	fetch := fakeFetch(nil, map[Month]error{start: want})
	_, _, err := SelectNext(start, map[string]bool{}, fetch)
	if !errors.Is(err, want) {
		t.Errorf("err = %v", err)
	}
}
