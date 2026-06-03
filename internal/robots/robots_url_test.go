//ff:func feature=robots type=helper control=sequence
//ff:what RobotsURL이 호스트로부터 https://<host>/robots.txt를 만드는지(빈 호스트 포함) 검증한다.

package robots

import "testing"

func TestRobotsURL(t *testing.T) {
	if got, want := RobotsURL("example.com"), "https://example.com/robots.txt"; got != want {
		t.Errorf("RobotsURL = %q, want %q", got, want)
	}
	if got, want := RobotsURL(""), "https:///robots.txt"; got != want {
		t.Errorf("RobotsURL(empty) = %q, want %q", got, want)
	}
}
