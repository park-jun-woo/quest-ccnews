//ff:func feature=extract type=helper control=sequence
//ff:what 단일 Article 객체의 필드들을 Fields로 매핑한다. headline→Title, author→Author, datePublished→PublishedAt, publisher→MediaName, inLanguage→Lang, articleBody→Body. 순수 함수.

package extract

// mapArticle maps a single Article object's fields to Fields.
func mapArticle(o map[string]any) Fields {
	return Fields{
		Title:       ldString(o["headline"]),
		Author:      ldName(o["author"]),
		PublishedAt: ldString(o["datePublished"]),
		MediaName:   ldName(o["publisher"]),
		Lang:        ldString(o["inLanguage"]),
		Body:        ldString(o["articleBody"]),
	}
}
