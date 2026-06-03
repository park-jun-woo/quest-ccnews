//ff:func feature=ingestion type=helper control=sequence
//ff:what ToUTF8가 UTF-8 입력에 멱등이고, Shift-JIS 선언 바이트를 올바른 UTF-8로 디코드하며, 디코드 불가 시 원본으로 폴백하는지 검증한다.

package ingest

import (
	"testing"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func TestToUTF8(t *testing.T) {
	// 멱등: UTF-8 입력은 그대로 통과.
	utf8in := []byte("한글과 日本語 mixed")
	out, _ := ToUTF8(utf8in, "text/html; charset=utf-8")
	if string(out) != string(utf8in) {
		t.Errorf("UTF-8 입력 멱등 실패: got %q", out)
	}

	// Shift-JIS 선언 바이트 → 올바른 일본어 UTF-8 복원.
	want := "こんにちは世界"
	sjis, _, err := transform.Bytes(japanese.ShiftJIS.NewEncoder(), []byte(want))
	if err != nil {
		t.Fatalf("fixture 인코딩 실패: %v", err)
	}
	got, _ := ToUTF8(sjis, "text/html; charset=Shift_JIS")
	if string(got) != want {
		t.Errorf("Shift-JIS 디코드 실패: got %q, want %q", got, want)
	}

	// 폴백: 빈 입력은 원본(빈 슬라이스) 그대로.
	if out, _ := ToUTF8([]byte{}, ""); len(out) != 0 {
		t.Errorf("빈 입력 폴백 실패: got %q", out)
	}
}
