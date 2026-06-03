//ff:func feature=cli type=helper control=sequence level=error
//ff:what runSubmit가 --url이 비면 "--url" 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunSubmit_MissingURL(t *testing.T) {
	resetSubmitFlags(t)
	submitURL = ""
	submitEvent6 = "ev.json"
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runSubmit(cmd, nil); err == nil || !strings.Contains(err.Error(), "--url") {
		t.Fatalf("want --url error, got %v", err)
	}
}
