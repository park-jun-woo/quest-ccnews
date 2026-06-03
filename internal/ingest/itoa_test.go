//ff:func feature=ingestion type=helper control=sequence
//ff:what 테스트 헬퍼. int를 10진 문자열로 변환한다(strconv.Itoa 래퍼).

package ingest

import "strconv"

func itoa(n int) string {
	return strconv.Itoa(n)
}
