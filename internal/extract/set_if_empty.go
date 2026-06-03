//ff:func feature=extract type=helper control=sequence
//ff:what *dst가 아직 비어있을 때만 v를 기록한다(필드별 첫 값 우선 보장). 순수 헬퍼.

package extract

// setIfEmpty writes v into *dst only when *dst is still empty.
func setIfEmpty(dst *string, v string) {
	if *dst == "" {
		*dst = v
	}
}
