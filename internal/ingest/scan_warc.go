//ff:func feature=ingestion type=helper control=iteration dimension=1 level=error
//ff:what 다운로드한 WARC 파일을 열어 response 레코드를 순회하며 TODO 기사를 만들어 콜백에 넘긴다(얇은 IO + 순수 ToArticle 호출).

package ingest

import (
	"io"
	"os"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/slyrz/warc"
)

// ScanWarc opens a downloaded WARC file, walks its records, and for every usable
// response record builds a TODO article (via the pure ToArticle) and hands it to
// emit. warcName is the lock/cursor key stored in each article's WARC locator.
//
// IO note: the slyrz reader transparently decompresses the record-gzip stream and
// does not expose true on-disk byte offsets, so Offset is a sequential record
// index within the WARC — sufficient as a {file,offset} locator for re-reading
// (the file plus record ordinal identifies the record). Body text is never stored.
func (c *Client) ScanWarc(warcPath, warcName string, emit func(*session.Article)) error {
	f, err := os.Open(warcPath)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := warc.NewReader(f)
	if err != nil {
		return err
	}
	defer r.Close()

	var ordinal int64
	for {
		rec, err := r.ReadRecord()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		rv := RecordView{
			Type:      rec.Header.Get("warc-type"),
			TargetURI: rec.Header.Get("warc-target-uri"),
			Offset:    ordinal,
		}
		ordinal++
		if a, ok := ToArticle(warcName, rv); ok {
			emit(a)
		}
	}
	return nil
}
