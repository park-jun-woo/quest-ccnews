//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what WARC 객체를 캐시 디렉터리로 받아 파일 경로를 돌려준다. 캐시에 동일 이름·동일 크기 파일이 있으면 본문 GET 없이 재사용한다. Content-Length와 받은 바이트가 다르면 부분 다운로드로 보고 거부·삭제한다(얇은 IO).

package ingest

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadWarc fetches the WARC at the given object path into the cache dir and
// returns the local file path. CC-NEWS WARCs are immutable content-addressed
// files (timestamp+sequence in the name), so a cache hit is reused instead of
// re-downloading ~1GB.
//
// Deterministic gate ("부분/불일치는 거부, 완전 캐시는 재사용"):
//   - If a file already exists at dest, a light HEAD probe (Range-GET fallback)
//     fetches the authoritative Content-Length. If the local size matches, the
//     body GET is skipped and the cached path is returned — the core saving.
//   - On any size mismatch (partial/corrupt cache) or an unknown remote size,
//     the full body is downloaded and dest is overwritten.
//   - The partial-download integrity gate is preserved: if the server
//     advertised a Content-Length that differs from the bytes written, the
//     partial file is removed and an error returned (Phase003).
func (c *Client) DownloadWarc(objectPath string) (string, error) {
	if err := os.MkdirAll(c.cacheDir, 0o755); err != nil {
		return "", err
	}
	dest := filepath.Join(c.cacheDir, WarcName(objectPath))

	// Cache-reuse gate: reuse a complete local file without fetching the body.
	if fi, err := os.Stat(dest); err == nil && fi.Mode().IsRegular() {
		if remote, ok := c.remoteSize(objectPath); ok && remote == fi.Size() {
			return dest, nil
		}
		// Unknown remote size or size mismatch (partial/corrupt): fall through
		// and re-download, overwriting dest below.
	}

	req, err := http.NewRequest(http.MethodGet, WarcURL(objectPath), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", c.userAgent)
	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download %s: status %d", objectPath, resp.StatusCode)
	}

	f, err := os.Create(dest)
	if err != nil {
		return "", err
	}
	n, copyErr := io.Copy(f, resp.Body)
	closeErr := f.Close()
	if copyErr != nil {
		os.Remove(dest)
		return "", copyErr
	}
	if closeErr != nil {
		os.Remove(dest)
		return "", closeErr
	}
	if resp.ContentLength >= 0 && n != resp.ContentLength {
		os.Remove(dest)
		return "", fmt.Errorf("download %s: partial (%d of %d bytes)", objectPath, n, resp.ContentLength)
	}
	return dest, nil
}
