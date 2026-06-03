//ff:func feature=output type=helper control=sequence level=error
//ff:what 레코드 1건을 out 파일에 append 모드로 한 줄(JSON+\n) 추가하는 IO. 디렉터리 없으면 생성. O(1) append.

package output

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Append marshals one Record and appends it as a single line (JSON + "\n") to
// the file at path, creating the parent directory if needed. It opens in append
// mode so each emit is an O(1) write — the JSONL file is never rewritten whole.
// This is the only IO in the package; rendering is pure (see Render).
func Append(path string, r *Record) error {
	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	line, err := json.Marshal(r)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(append(line, '\n')); err != nil {
		return err
	}
	return nil
}
