//ff:func feature=ingestion type=helper control=sequence
//ff:what ParsePaths가 빈 줄·공백뿐인 본문에 대해 nil을 반환하는지 검증한다.

package ingest

import "testing"

func TestParsePathsEmpty(t *testing.T) {
	if got := ParsePaths("\n  \n\t\n"); got != nil {
		t.Errorf("ParsePaths(blank) = %v, want nil", got)
	}
}
