//ff:func feature=ingestion type=helper control=sequence dimension=1 level=error
//ff:what raw 바이트와 Content-Type을 받아 x/net charset.DetermineEncoding 한 호출로 인코딩을 판정하고 transform으로 UTF-8 디코드해 반환한다. UTF-8 입력은 멱등, 디코드 실패 시 원본으로 안전 폴백. 순수 함수(네트워크/파일 IO 없음).

package ingest

import (
	"bytes"
	"io"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

// ToUTF8 decodes raw response bytes to UTF-8. It delegates encoding detection to
// golang.org/x/net/html/charset.DetermineEncoding in a single call — that helper
// already inspects, in priority order, the declared Content-Type charset, a BOM,
// the HTML <meta> prescan, and a heuristic, so we do not re-implement any of those
// steps (Phase008 §인코딩 결정). The detected encoding's decoder is run through a
// transform.Reader to produce UTF-8 bytes.
//
// UTF-8 input is idempotent: DetermineEncoding returns "utf-8", whose decoder is a
// pass-through. On decode failure (or empty result) it falls back to the original
// bytes assumed to be UTF-8 (Phase008 실패 정책 — 과도한 SKIP 회피). Pure: no IO.
func ToUTF8(raw []byte, contentType string) (out []byte, encName string) {
	enc, name, _ := charset.DetermineEncoding(raw, contentType)
	decoded, err := io.ReadAll(transform.NewReader(bytes.NewReader(raw), enc.NewDecoder()))
	if err != nil || len(decoded) == 0 {
		return raw, name
	}
	return decoded, name
}
