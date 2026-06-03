//ff:func feature=output type=helper control=sequence level=error
//ff:what Append가 디렉터리 자동 생성·한 줄 유효 JSON append·중복 호출 누적·bare 파일명·MkdirAll/OpenFile 오류를 올바르게 처리하는지 검증한다.

package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppend(t *testing.T) {
	t.Run("creates dir and writes line", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "nested", "deep", "out.jsonl")

		rec := &Record{URL: "https://e.com/a", Host: "e.com", Status: "PASS"}
		if err := Append(path, rec); err != nil {
			t.Fatalf("Append() error = %v", err)
		}
		if _, err := os.Stat(filepath.Dir(path)); err != nil {
			t.Fatalf("parent dir not created: %v", err)
		}

		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("ReadFile() error = %v", err)
		}
		if !strings.HasSuffix(string(data), "\n") {
			t.Errorf("line not terminated with newline: %q", string(data))
		}

		var got Record
		if err := json.Unmarshal([]byte(strings.TrimSpace(string(data))), &got); err != nil {
			t.Fatalf("line is not valid JSON: %v", err)
		}
		if got.URL != rec.URL || got.Host != rec.Host || got.Status != rec.Status {
			t.Errorf("round-trip = %+v, want %+v", got, rec)
		}
	})

	t.Run("accumulates lines", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "out.jsonl")

		for i, st := range []string{"PASS", "REVIEW", "BLOCKED"} {
			if err := Append(path, &Record{URL: "u", Status: st}); err != nil {
				t.Fatalf("Append() #%d error = %v", i, err)
			}
		}

		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("ReadFile() error = %v", err)
		}
		lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
		if len(lines) != 3 {
			t.Fatalf("got %d lines, want 3: %q", len(lines), string(data))
		}
		for _, l := range lines {
			var r Record
			if err := json.Unmarshal([]byte(l), &r); err != nil {
				t.Errorf("line %q not valid JSON: %v", l, err)
			}
		}
	})

	t.Run("bare filename", func(t *testing.T) {
		dir := t.TempDir()
		orig, err := os.Getwd()
		if err != nil {
			t.Fatalf("Getwd: %v", err)
		}
		if err := os.Chdir(dir); err != nil {
			t.Fatalf("Chdir: %v", err)
		}
		defer os.Chdir(orig)

		if err := Append("bare.jsonl", &Record{URL: "u", Status: "PASS"}); err != nil {
			t.Fatalf("Append() error = %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "bare.jsonl")); err != nil {
			t.Fatalf("file not written: %v", err)
		}
	})

	t.Run("mkdir error", func(t *testing.T) {
		dir := t.TempDir()
		blocker := filepath.Join(dir, "blocker")
		if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup WriteFile: %v", err)
		}
		if err := Append(filepath.Join(blocker, "child", "out.jsonl"), &Record{}); err == nil {
			t.Errorf("expected error appending under a file path, got nil")
		}
	})

	t.Run("open error", func(t *testing.T) {
		dir := t.TempDir()
		target := filepath.Join(dir, "iamadir")
		if err := os.Mkdir(target, 0o755); err != nil {
			t.Fatalf("setup Mkdir: %v", err)
		}
		if err := Append(target, &Record{URL: "u"}); err == nil {
			t.Errorf("expected error opening a directory for write, got nil")
		}
	})
}
