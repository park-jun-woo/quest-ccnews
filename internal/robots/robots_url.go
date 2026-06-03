//ff:func feature=robots type=helper control=sequence
//ff:what 호스트로부터 https://<host>/robots.txt URL을 만든다. 순수 함수.

package robots

// RobotsURL builds the canonical robots.txt URL for a host. We always probe
// over https (Phase004: hosts are public news sites). Pure — no IO.
func RobotsURL(host string) string {
	return "https://" + host + "/robots.txt"
}
