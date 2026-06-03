//ff:func feature=ingestion type=helper control=iteration dimension=1
//ff:what TracksFor가 forward/backward/both/기타 플래그를 올바른 트랙 목록으로 매핑하는지 검증한다.

package ingest

import (
	"reflect"
	"testing"
)

func TestTracksFor(t *testing.T) {
	cases := []struct {
		flag string
		want []Track
	}{
		{"forward", []Track{Forward}},
		{"backward", []Track{Backward}},
		{"both", []Track{Forward, Backward}},
		{"", []Track{Forward, Backward}},
		{"anything", []Track{Forward, Backward}},
	}
	for _, c := range cases {
		if got := TracksFor(c.flag); !reflect.DeepEqual(got, c.want) {
			t.Errorf("TracksFor(%q) = %v, want %v", c.flag, got, c.want)
		}
	}
}
