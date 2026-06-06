//ff:func feature=event6 type=helper control=sequence
//ff:what (*jsonScanState).feedInString 분기 테스트. 백슬래시가 다음 바이트를 이스케이프, 이스케이프된 따옴표는 종료 안 함, 이스케이프 안 된 따옴표는 문자열 종료.

package ccnewsquest

import "testing"

func TestJSONScanStateFeedInString(t *testing.T) {
	t.Run("backslash arms escape, escaped quote stays in string", func(t *testing.T) {
		st := jsonScanState{inString: true}
		st.feedInString('\\') // arm escape
		if !st.escaped {
			t.Fatal("backslash should arm escape")
		}
		st.feedInString('"') // escaped quote → consumed, still in string
		if st.escaped {
			t.Fatal("escape should be cleared after the escaped byte")
		}
		if !st.inString {
			t.Fatal("escaped quote must not end the string")
		}
	})

	t.Run("unescaped quote ends string", func(t *testing.T) {
		st := jsonScanState{inString: true}
		st.feedInString('x') // ordinary byte
		if !st.inString {
			t.Fatal("ordinary byte must not end the string")
		}
		st.feedInString('"')
		if st.inString {
			t.Fatal("unescaped quote should end the string")
		}
	})
}
