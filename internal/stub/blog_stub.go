package stub

import "dot-golang/internal/domain"

const (
	StubCacheNews = `[
		{
			"id": 1,
			"title": "test lorem",
			"content": "test content lorem",
		},
		{
			"id": 2,
			"title": "test ipsum",
			"content": "test content ipsum",
		}
	]`
)

var StubNews = []domain.News{
	{
		Id:            1,
		Title:         "test lorem",
		Content:       "test content lorem",
		StatusContent: domain.PUBLISHED,
	},
	{
		Id:            2,
		Title:         "test ipsum",
		Content:       "test content ipsum",
		StatusContent: domain.PUBLISHED,
	},
}
