package abstraction

import "dot-golang/internal/domain"

type BlogService interface {
	GetNews(page int) ([]domain.News, error)
	GetNewsDrafted(page int) ([]domain.News, error)
	GetNewsPublished(page int) ([]domain.News, error)
	GetNewsDeleted(page int) ([]domain.News, error)
	PostNews(new domain.News) error
	GetSingleNews(id int64) (domain.News, error)
	PutSingleNews(new domain.News) error
	DeleteSingleNews(id int64) error
	PatchSingleNewsStatusContent(id int64, statusContent domain.StatusContent) error
	GetTopics(page int) ([]domain.Topics, error)
	GetSingleTopics(id int64) (domain.Topics, error)
	PostTopics(topic domain.Topics) error
	PatchSingleTopicsCategoryName(idTopic int64, categoryName string) error
	GetSingleTopicsNews(idTopic int64, page int) ([]domain.News, error)
	PostRefNewsTopics(domain.RefNewsTopics) error
	DeleteRefNewsTopics(domain.RefNewsTopics) error
}
