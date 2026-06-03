//ff:func feature=article type=helper control=iteration dimension=1
//ff:what URL로 기사 퀘스트를 찾는다. 없으면 에러를 반환한다.

package session

import "fmt"

// Find: locate an article quest by its URL.
func (s *Session) Find(url string) (*Article, error) {
	for _, a := range s.Articles {
		if a.URL == url {
			return a, nil
		}
	}
	return nil, fmt.Errorf("article not found in session: %q", url)
}
