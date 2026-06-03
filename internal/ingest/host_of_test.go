//ff:func feature=ingestion type=helper control=sequence
//ff:what HostOf가 URL에서 소문자 host(포트 제외)를 뽑는지 검증한다.

package ingest

import "testing"

func TestHostOfLowercasesAndStripsPort(t *testing.T) {
	if h, ok := HostOf("https://www.Example.com/path?q=1"); !ok || h != "www.example.com" {
		t.Errorf("HostOf = %q,%v want www.example.com,true", h, ok)
	}
	if h, ok := HostOf("http://news.site.org:8080/a"); !ok || h != "news.site.org" {
		t.Errorf("HostOf(port) = %q,%v want news.site.org,true", h, ok)
	}
	if h, ok := HostOf("https://EXAMPLE.COM"); !ok || h != "example.com" {
		t.Errorf("HostOf(upper) = %q,%v want example.com,true", h, ok)
	}
}
