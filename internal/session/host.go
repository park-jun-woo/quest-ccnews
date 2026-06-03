//ff:type feature=host type=model
//ff:what 호스트별 캐시. robots 상태와 매체 정보, CC 라이선스를 담는다. 추출 템플릿은 두지 않는다(구조화 데이터만 신뢰).

package session

// Host: per-host cache holding robots state and media info. No extraction
// templates: only structured data is trusted.
type Host struct {
	MediaName string   `json:"media_name,omitempty"` // JSON-LD publisher.name / og:site_name
	SiteURL   string   `json:"site_url,omitempty"`   // media site address (user-specified field)
	Robots    *Robots  `json:"robots,omitempty"`
	License   *License `json:"license"` // CC license: {type,url,source} | null
}
