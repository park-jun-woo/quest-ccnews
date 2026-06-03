//ff:func feature=robots type=helper control=sequence
//ff:what NewClient가 15초 타임아웃 HTTP 클라이언트와 UA 헤더·product token을 채운 Client를 만드는지 검증한다.

package robots

import (
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c := NewClient("parkjunwoo-quest/0.1 (+https://example.com)")
	if c == nil {
		t.Fatal("NewClient returned nil")
	}
	if c.userAgent != "parkjunwoo-quest/0.1 (+https://example.com)" {
		t.Errorf("userAgent = %q", c.userAgent)
	}
	if c.productToken != "parkjunwoo-quest" {
		t.Errorf("productToken = %q, want parkjunwoo-quest", c.productToken)
	}
	if c.http == nil {
		t.Fatal("http client is nil")
	}
	if c.http.Timeout != 15*time.Second {
		t.Errorf("timeout = %v, want 15s", c.http.Timeout)
	}
}
