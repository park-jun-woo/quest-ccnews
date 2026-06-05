//ff:func feature=gate type=helper control=iteration dimension=1 level=error
//ff:what Definition.Seed. 위치인자(기사 URL 목록)에서 Key=URL, Payload=&session.Article{URL:…}인 TODO 아이템들을 시드한다. 최소 구현 — 실제 투트랙 WARC 인제스천(다운로드·레코드·커서)은 Phase013 run 명령 소관. 빈/공백 URL은 버린다.

package ccnewsquest

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

// Seed creates one TODO item per article URL positional arg, with Key=URL and
// Payload=&session.Article{URL: ...}. This is the minimal seed: the real two-track
// CC-NEWS ingestion (WARC download, record walk, cursor advance, robots BLOCKED
// seeding) is Phase013's `run` command, which will populate Payload.WARC. Blank
// args are dropped.
func (ccnewsDef) Seed(args []string) ([]*quest.Item, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("seed: 기사 URL이 필요합니다 (예: ccnews scan https://example.com/a https://example.com/b)")
	}
	items := make([]*quest.Item, 0, len(args))
	for _, arg := range args {
		url := strings.TrimSpace(arg)
		if url == "" {
			continue
		}
		items = append(items, &quest.Item{
			Key:     url,
			State:   quest.TODO,
			Payload: &session.Article{URL: url},
		})
	}
	return items, nil
}
