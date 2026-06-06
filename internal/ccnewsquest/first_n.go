//ff:func feature=event6 type=helper control=sequence
//ff:what firstN — 바이트열 앞 n바이트를 문자열로(짧으면 그대로). 디코드 실패 verdict의 Fact.Actual에 원시 출력 앞부분을 싣는 용도.

package ccnewsquest

// firstN returns the first n bytes of b as a string (or all of b when shorter). It
// is used to put a short prefix of an undecodable submission into the FAIL verdict's
// Fact.Actual so the agent sees what it produced without dumping the whole blob.
func firstN(b []byte, n int) string {
	if len(b) <= n {
		return string(b)
	}
	return string(b[:n])
}
