//ff:func feature=robots type=helper control=sequence
//ff:what Parse가 주석 줄과 알 수 없는 필드(Sitemap)를 무시하는지 검증한다.

package robots

import "testing"

func TestParseIgnoresCommentsAndUnknown(t *testing.T) {
	content := []byte("# comment\nUser-agent: *\nSitemap: https://x/s.xml\nDisallow: /a\n")
	rs := Parse(content)
	if len(rs.Groups) != 1 || len(rs.Groups[0].Rules) != 1 {
		t.Fatalf("unexpected ruleset: %+v", rs)
	}
}
