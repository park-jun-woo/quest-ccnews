//ff:func feature=ingestion type=helper control=sequence
//ff:what HostOf가 파싱 실패·host 부재 URL에 대해 ok=false를 반환하는지 검증한다.

package ingest

import "testing"

func TestHostOfNoHost(t *testing.T) {
	if h, ok := HostOf("://bad-url"); ok || h != "" {
		t.Errorf("HostOf(bad) = %q,%v want '',false", h, ok)
	}
	if h, ok := HostOf("/just/a/path"); ok || h != "" {
		t.Errorf("HostOf(path) = %q,%v want '',false", h, ok)
	}
	if h, ok := HostOf(""); ok || h != "" {
		t.Errorf("HostOf(empty) = %q,%v want '',false", h, ok)
	}
}
