//ff:func feature=output type=helper control=sequence
//ff:what bool 값의 주소를 돌려주는 헬퍼. crawl_allowed를 항상 present로 직렬화하기 위함(false도 표현). 순수.

package output

// boolPtr returns the address of b so a bool field is always present in JSON
// (even when false) without colliding with omitempty.
func boolPtr(b bool) *bool { return &b }
