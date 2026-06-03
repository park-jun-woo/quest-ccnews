//ff:func feature=ingestion type=helper control=sequence
//ff:what NewClient가 user-agent, cacheDir, 기본 30분 타임아웃 HTTP 클라이언트를 세팅하는지 검증한다.

package ingest

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c := NewClient("ua/1.0", "/tmp/cache")
	if c == nil {
		t.Fatal("NewClient returned nil")
	}
	if c.userAgent != "ua/1.0" {
		t.Errorf("userAgent = %q", c.userAgent)
	}
	if c.cacheDir != "/tmp/cache" {
		t.Errorf("cacheDir = %q", c.cacheDir)
	}
	if c.http == nil {
		t.Fatal("http client is nil")
	}
	if c.http.Timeout != 30*time.Minute {
		t.Errorf("timeout = %v, want 30m", c.http.Timeout)
	}
}
