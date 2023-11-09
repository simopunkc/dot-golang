package abstraction

import "dot-golang/internal/domain"

type BlogRepository interface {
	GetNews(limit int, offset int) ([]domain.News, error)
	GetNewsDrafted(limit int, offset int) ([]domain.News, error)
	GetNewsPublished(limit int, offset int) ([]domain.News, error)
	GetNewsDeleted(limit int, offset int) ([]domain.News, error)
	PostNews(new domain.News) error
	GetSingleNews(id int64) (domain.News, error)
	PutSingleNews(new domain.News) error
	DeleteSingleNews(id int64) error
	PatchSingleNewsStatusContent(id int64, statusContent domain.StatusContent) error
	GetTopics(limit int, offset int) ([]domain.Topics, error)
	GetSingleTopics(id int64) (domain.Topics, error)
	PostTopics(topics domain.Topics) error
	PatchSingleTopicsCategoryName(idTopic int64, categoryName string) error
	GetSingleTopicsNews(idTopic int64, limit int, offset int) ([]domain.News, error)
	PostRefNewsTopics(refNewsTopics domain.RefNewsTopics) error
	DeleteRefNewsTopics(refNewsTopics domain.RefNewsTopics) error
}
