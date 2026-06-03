//ff:func feature=ingestion type=helper control=iteration dimension=1 level=error
//ff:what WARC 로케이터{file,offset}로 캐시의 .warc.gz를 다시 열어 그 레코드의 HTTP 응답 본문(HTML 바이트)을 돌려준다. 본문 미저장 원칙에 따라 앵커 검증마다 WARC에서 재독하는 얇은 IO.

package ingest

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/quest-ccnews/internal/session"
	"github.com/slyrz/warc"
)

// ReadBody re-reads the HTML body for an article from its WARC locator. Bodies
// are never stored in the session (the WARC is the source of truth — Phase002
// 결정 1), so the anchor gate (Phase006) re-reads on every check. It opens
// cacheDir/<loc.File>, walks records using the same ordinal counter as ScanWarc,
// finds the record at loc.Offset, parses its HTTP response, and returns the
// response body bytes (the HTML). Thin IO; the HTML→text/event6 logic is pure
// elsewhere (extract.Parse, anchor.Gate).
func (c *Client) ReadBody(loc *session.WARCLoc) ([]byte, error) {
	if loc == nil {
		return nil, fmt.Errorf("nil WARC locator")
	}
	path := filepath.Join(c.cacheDir, loc.File)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r, err := warc.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var ordinal int64
	for {
		rec, err := r.ReadRecord()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if ordinal == loc.Offset {
			return httpResponseBody(rec.Content)
		}
		ordinal++
	}
	return nil, fmt.Errorf("record offset %d not found in %s", loc.Offset, loc.File)
}

// httpResponseBody parses a WARC response record's content (a raw HTTP response:
// status line + headers + body) and returns the body bytes. CC-NEWS response
// records wrap the origin's HTTP response, so the HTML payload follows the HTTP
// header block.
func httpResponseBody(content io.Reader) ([]byte, error) {
	resp, err := http.ReadResponse(bufio.NewReader(content), nil)
	if err != nil {
		return nil, fmt.Errorf("parse WARC HTTP response: %w", err)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
