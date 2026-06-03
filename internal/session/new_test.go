//ff:func feature=session type=helper control=sequence
//ff:what New가 버전·user agent·source를 세팅하고 빈 호스트 맵/기사 목록을 만드는지 검증한다.

package session

import "testing"

func TestNew(t *testing.T) {
	s := New("ccnews-bot/1.0", "cc-news")
	if s == nil {
		t.Fatal("New returned nil")
	}
	if s.Version != 1 {
		t.Errorf("Version = %d, want 1", s.Version)
	}
	if s.UserAgent != "ccnews-bot/1.0" {
		t.Errorf("UserAgent = %q, want %q", s.UserAgent, "ccnews-bot/1.0")
	}
	if s.Ingestion.Source != "cc-news" {
		t.Errorf("Ingestion.Source = %q, want %q", s.Ingestion.Source, "cc-news")
	}
	if s.Ingestion.ProcessedWarcs == nil || len(s.Ingestion.ProcessedWarcs) != 0 {
		t.Errorf("ProcessedWarcs = %v, want empty non-nil", s.Ingestion.ProcessedWarcs)
	}
	if s.Hosts == nil || len(s.Hosts) != 0 {
		t.Errorf("Hosts = %v, want empty non-nil", s.Hosts)
	}
	if s.Articles == nil || len(s.Articles) != 0 {
		t.Errorf("Articles = %v, want empty non-nil", s.Articles)
	}
}
