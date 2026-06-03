//ff:func feature=ingestion type=helper control=sequence
//ff:what 테스트 헬퍼. month→paths 맵과 실패 맵으로부터 fetchPaths 콜백을 만든다(IO 없는 주입용).

package ingest

// fakeFetch builds a fetchPaths callback from a month→paths map.
func fakeFetch(byMonth map[Month][]string, fail map[Month]error) func(Month) ([]string, error) {
	return func(m Month) ([]string, error) {
		if fail != nil {
			if err, ok := fail[m]; ok {
				return nil, err
			}
		}
		return byMonth[m], nil
	}
}
