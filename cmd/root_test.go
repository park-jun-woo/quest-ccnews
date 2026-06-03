//ff:func feature=cli type=command control=sequence
//ff:what Execute()가 인자 없이 호출될 때 os.Exit(1) 분기 없이 정상 종료하는지 검증한다(root.go 커버).

package cmd

import (
	"bytes"
	"testing"
)

// TestExecute covers the normal (no-error) path of Execute. The os.Exit(1)
// error branch is intentionally not exercised (best-effort coverage).
func TestExecute(t *testing.T) {
	// Drive the root command with no args so it prints usage and returns nil,
	// avoiding the os.Exit(1) branch.
	rootCmd.SetArgs([]string{})
	var out bytes.Buffer
	rootCmd.SetOut(&out)
	rootCmd.SetErr(&out)

	// Execute should not call os.Exit since rootCmd.Execute returns nil here.
	Execute()
}
