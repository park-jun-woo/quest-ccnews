//ff:func feature=session type=helper control=sequence
//ff:what Save와 Load의 왕복으로 세션이 손실 없이 직렬화·복원되는지 검증한다.

package session

import (
	"path/filepath"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "session.json")

	orig := New("agent/1.0", "cc-news")
	orig.Articles = append(orig.Articles, &Article{
		URL:   "https://example.com/x",
		Host:  "example.com",
		State: TODO,
	})
	orig.Hosts["example.com"] = &Host{MediaName: "Example News"}

	if err := orig.Save(path); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if loaded.UserAgent != orig.UserAgent {
		t.Errorf("UserAgent = %q, want %q", loaded.UserAgent, orig.UserAgent)
	}
	if loaded.Ingestion.Source != orig.Ingestion.Source {
		t.Errorf("Source = %q, want %q", loaded.Ingestion.Source, orig.Ingestion.Source)
	}
	if len(loaded.Articles) != 1 || loaded.Articles[0].URL != "https://example.com/x" {
		t.Errorf("Articles = %+v, want one article with URL example.com/x", loaded.Articles)
	}
	if h, ok := loaded.Hosts["example.com"]; !ok || h.MediaName != "Example News" {
		t.Errorf("Hosts[example.com] = %+v, want MediaName Example News", loaded.Hosts["example.com"])
	}
}
