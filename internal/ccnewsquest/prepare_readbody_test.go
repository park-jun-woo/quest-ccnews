//ff:func feature=gate type=helper control=sequence
//ff:what Prepare WARC ReadBody 실패 단위테스트. 존재하지 않는 캐시 파일을 가리키는 로케이터로 ingest.Client.ReadBody의 os.Open 실패 분기를 결정적으로 탄다(네트워크·실 WARC 불요).

package ccnewsquest

import (
	"testing"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/park-jun-woo/reins/pkg/quest"
)

func TestPrepareReadBodyError(t *testing.T) {
	// A locator pointing at a file that does not exist in cacheDir makes
	// ingest.Client.ReadBody fail at os.Open, exercising Prepare's ReadBody
	// error branch deterministically (no network, no real WARC).
	d := Def("ua", t.TempDir())
	it := &quest.Item{Key: "https://x/a"}
	if err := it.SetPayload(&session.Article{
		URL:  "https://x/a",
		WARC: &session.WARCLoc{File: "does-not-exist.warc.gz", Offset: 0},
	}); err != nil {
		t.Fatal(err)
	}
	_, _, err := d.Prepare(it, []byte(`{"who":{"value":"Alice","anchors":["Alice"]}}`))
	if err == nil {
		t.Fatal("want error when ReadBody cannot open the WARC file")
	}
}
