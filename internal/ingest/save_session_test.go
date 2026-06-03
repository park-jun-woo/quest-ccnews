//ff:func feature=ingestion type=helper control=sequence
//ff:what saveSession이 Save 콜백 미설정(nil)·성공·실패 세 분기를 올바르게 처리하는지 검증한다.

package ingest

import (
	"errors"
	"strings"
	"testing"
)

func TestSaveSession(t *testing.T) {
	// 콜백 미설정 → nil
	if err := saveSession(RunOptions{Save: nil}); err != nil {
		t.Fatalf("nil Save should return nil, got %v", err)
	}

	// 콜백 성공 → nil
	if err := saveSession(RunOptions{Save: func() error { return nil }}); err != nil {
		t.Fatalf("successful Save should return nil, got %v", err)
	}

	// 콜백 실패 → 래핑된 에러
	err := saveSession(RunOptions{Save: func() error { return errors.New("boom") }})
	if err == nil {
		t.Fatal("failing Save should return error")
	}
	if !strings.Contains(err.Error(), "save session") || !strings.Contains(err.Error(), "boom") {
		t.Fatalf("error should wrap cause, got %v", err)
	}
}
