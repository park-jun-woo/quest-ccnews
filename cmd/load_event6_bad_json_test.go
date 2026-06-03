//ff:func feature=cli type=helper control=sequence level=error
//ff:what loadEvent6가 잘못된 JSON에서 "parse event6 JSON" 에러를 반환하는지 검증한다.

package cmd

import (
	"strings"
	"testing"
)

func TestLoadEvent6_BadJSON(t *testing.T) {
	_, err := loadEvent6("-", strings.NewReader("{not json"))
	if err == nil {
		t.Fatal("want parse error")
	}
	if !strings.Contains(err.Error(), "parse event6 JSON") {
		t.Errorf("err = %v", err)
	}
}
