//ff:func feature=robots type=helper control=iteration dimension=1
//ff:what elbalad.news 류 robots.txt(빈 Disallow + 둘째 * 레코드의 실제 경로 규칙) 픽스처로 /6993730 허용·실제 Disallow 경로 차단을 검증한다(실버그 재현→해소).

package robots

import "testing"

// elbaladFixture mirrors the shape of www.elbalad.news/robots.txt that produced
// the Phase013 false BLOCKED: a first "User-agent: *" record carrying only an
// empty Disallow (allow-all idiom), then a second "User-agent: *" record with
// the real path restrictions. D1 drops the empty Disallow; D2 merges both
// records so the real restrictions still apply.
const elbaladFixture = `User-agent: *
Disallow:

User-agent: *
Disallow: /reporter/
Disallow: /search
Disallow: /tags/
Crawl-delay: 10
`

func TestEvaluateElbaladFixture(t *testing.T) {
	rs := Parse([]byte(elbaladFixture))

	if d := Evaluate(rs, "parkjunwoo-quest", "/6993730"); !d.Allowed {
		t.Errorf("/6993730 should be Allowed (real bug), got %+v", d)
	}

	blocked := []string{"/reporter/123", "/search?q=x", "/tags/news"}
	for _, p := range blocked {
		if d := Evaluate(rs, "parkjunwoo-quest", p); d.Allowed {
			t.Errorf("path %q should be blocked, got %+v", p, d)
		}
	}
}
