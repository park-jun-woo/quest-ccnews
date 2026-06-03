//ff:type feature=output type=model
//ff:what session.License를 출력용으로 미러링한 CC 라이선스 필드(type/url/source).

package output

// FieldLicense mirrors session.License for output (CC license type/url/source).
type FieldLicense struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Source string `json:"source"`
}
