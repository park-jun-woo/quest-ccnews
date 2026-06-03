//ff:func feature=ingestion type=helper control=sequence
//ff:what ProcessedSet이 슬라이스를 조회용 집합으로 변환하고 nil도 빈 맵으로 다루는지 검증한다.

package ingest

import "testing"

func TestProcessedSet(t *testing.T) {
	set := ProcessedSet([]string{"a", "b", "a"})
	if !set["a"] || !set["b"] {
		t.Errorf("set missing entries: %v", set)
	}
	if set["c"] {
		t.Error("set should not contain c")
	}
	if got := ProcessedSet(nil); len(got) != 0 {
		t.Errorf("ProcessedSet(nil) len = %d", len(got))
	}
}
