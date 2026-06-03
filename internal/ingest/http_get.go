//ff:func feature=ingestion type=helper control=sequence level=error
//ff:what user-agent를 붙여 URL로 GET 요청하고 200이 아니면 에러를 반환한다. 응답 본문 reader를 그대로 넘긴다(얇은 IO).

package ingest

import (
	"fmt"
	"io"
	"net/http"
)

// httpGet performs a GET with the configured user-agent and returns the response
// body reader on HTTP 200. The caller must Close the returned ReadCloser. Thin IO
// wrapper — no parsing here.
func httpGet(client *http.Client, url, userAgent string) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("GET %s: status %d", url, resp.StatusCode)
	}
	return resp.Body, nil
}
