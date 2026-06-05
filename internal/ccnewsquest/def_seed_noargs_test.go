//ff:func feature=gate type=helper control=sequence
//ff:what Def().Seed() 빈 인자 가드. URL이 하나도 없으면 에러를 내는지 검증한다.

package ccnewsquest

import "testing"

func TestSeedNoArgs(t *testing.T) {
	if _, err := Def("ua", "cache").Seed(nil); err == nil {
		t.Fatal("want error for empty args")
	}
}
