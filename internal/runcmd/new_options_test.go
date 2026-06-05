//ff:func feature=ingestion type=command control=sequence
//ff:what newOptions 테스트 헬퍼. 상속된 --session 플래그가 sessionPath를 가리키는 cobra 명령에 묶인 run options(track=forward, 출력 버퍼)를 만든다. 네트워크는 ingestRun 패키지 변수 교체로 회피한다.

package runcmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

// newOptions builds run options bound to a cobra command whose inherited
// --session flag points at sessionPath. Network is avoided by overriding the
// ingestRun package var per test.
func newOptions(t *testing.T, sessionPath string) (*options, *cobra.Command) {
	t.Helper()
	cmd := &cobra.Command{Use: "run"}
	var sess string
	cmd.Flags().StringVar(&sess, "session", sessionPath, "")
	cmd.SetOut(&bytes.Buffer{})
	o := &options{userAgent: "ua/0.1", cmd: cmd, track: "forward"}
	return o, cmd
}
