//ff:type feature=host type=model
//ff:what 호스트가 선언한 CC 라이선스(type/url/source).

package session

// License: the host's declared CC license.
type License struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Source string `json:"source"`
}
