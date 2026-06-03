//ff:func feature=anchor type=helper control=sequence
//ff:what 앵커 매칭 정규화. 유니코드 공백 런을 단일 스페이스로 접고 양끝을 트림한다. 원문·앵커 양쪽에 동일 적용(매핑 추론 아님). 순수 함수.

package anchor

import "strings"

// normalize collapses every run of Unicode whitespace to a single space and
// trims the ends. It is applied identically to the source body text and to each
// anchor so matching compares like-for-like surface forms — never an inferred
// mapping (Phase006 열린 결정: 정규화는 원문·앵커 양쪽 동일 적용). strings.Fields
// splits on Unicode whitespace, matching extract.stripHTML's own normalization.
func normalize(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
