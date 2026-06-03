//ff:func feature=cli type=helper control=sequence
//ff:what 테스트 헬퍼. submit 전역 플래그(submitURL/submitEvent6/sessionPath)를 테스트 종료 시 기본값으로 되돌린다.

package cmd

import "testing"

func resetSubmitFlags(t *testing.T) {
	t.Cleanup(func() {
		submitURL = ""
		submitEvent6 = ""
		sessionPath = "session.json"
	})
}
