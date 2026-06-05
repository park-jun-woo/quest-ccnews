//ff:func feature=ingestion type=command control=sequence
//ff:what flushOptions 테스트 헬퍼. 상속된 --session 플래그가 sessionPath를 가리키는 cobra 명령에 묶인 run options를 만들어 flush 테스트가 세션 경로를 제어하게 한다.

package runcmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func flushOptions(t *testing.T, sessionPath string) *options {
	t.Helper()
	cmd := &cobra.Command{Use: "run"}
	var sess string
	cmd.Flags().StringVar(&sess, "session", sessionPath, "")
	return &options{userAgent: "ua/0.1", cmd: cmd}
}
