//ff:func feature=cli type=helper control=sequence level=error
//ff:what runSubmit가 --event6가 비면 "--event6" 에러를 반환하는지 검증한다.

package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRunSubmit_MissingEvent6(t *testing.T) {
	resetSubmitFlags(t)
	submitURL = "https://example.com/a"
	submitEvent6 = ""
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	if err := runSubmit(cmd, nil); err == nil || !strings.Contains(err.Error(), "--event6") {
		t.Fatalf("want --event6 error, got %v", err)
	}
}
