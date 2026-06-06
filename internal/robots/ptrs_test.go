//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what 그룹 포인터 슬라이스를 비교 가능한 문자열로 직렬화하는 테스트 헬퍼. classifyGroup 누적 결과의 동일성 단언에 쓴다.

package robots

import "fmt"

func ptrs(gs []*Group) string {
	out := ""
	for _, g := range gs {
		out += fmt.Sprintf("%p;", g)
	}
	return out
}
