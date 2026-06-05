//ff:func feature=ingestion type=command control=sequence
//ff:what sessionPath 단위테스트. cmd가 nil이거나 session 플래그가 없으면 기본 "session.json", 상속된 session 플래그 값이 있으면 그 값을 돌려주는지 검증한다.

package runcmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestSessionPath(t *testing.T) {
	t.Run("nil cmd falls back to default", func(t *testing.T) {
		o := &options{}
		if got := o.sessionPath(); got != "session.json" {
			t.Errorf("sessionPath() = %q, want session.json", got)
		}
	})

	t.Run("missing session flag falls back to default", func(t *testing.T) {
		o := &options{cmd: &cobra.Command{Use: "run"}}
		if got := o.sessionPath(); got != "session.json" {
			t.Errorf("sessionPath() = %q, want session.json", got)
		}
	})

	t.Run("inherited session flag value is returned", func(t *testing.T) {
		cmd := &cobra.Command{Use: "run"}
		var sess string
		cmd.Flags().StringVar(&sess, "session", "custom.json", "")
		o := &options{cmd: cmd}
		if got := o.sessionPath(); got != "custom.json" {
			t.Errorf("sessionPath() = %q, want custom.json", got)
		}
	})
}
