package service

import (
	"dot-golang/internal/abstraction"
	"dot-golang/internal/domain"
)

//go:generate moq -out blog_service_mock_test.go -pkg service ../abstraction BlogCache BlogRepository BlogEvent BlogUtil

type BlogService struct {
	blogCache abstraction.BlogCache
	blogRepo  abstraction.BlogRepository
	blogEvent abstraction.BlogEvent
	blogUtil  abstraction.BlogUtil
}

func NewBlogService(blogCache abstraction.BlogCache, blogRepo abstraction.BlogRepository, blogEvent abstraction.BlogEvent, blogUtil abstraction.BlogUtil) *BlogService {
	return &BlogService{blogCache, blogRepo, blogEvent, blogUtil}
}

const (
	PREFIX_KEYCACHE_NEWS               = "GetNews"
	PREFIX_KEYCACHE_NEWS_DRAFTED       = "GetNewsDrafted"
	PREFIX_KEYCACHE_NEWS_PUBLISHED     = "GetNewsPublished"
	PREFIX_KEYCACHE_NEWS_DELETED       = "GetNewsDeleted"
	PREFIX_KEYCACHE_SINGLE_NEWS        = "GetSingleNews"
	PREFIX_KEYCACHE_SINGLE_TOPICS_NEWS = "GetSingleTopicsNews"
	PREFIX_KEYCACHE_TOPICS             = "GetTopics"
	PREFIX_KEYCACHE_SINGLE_TOPICS      = "GetSingleTopics"
)

func (bs BlogService) GetNews(page int) ([]domain.News, error) {
	limit, offset := bs.blogUtil.GetLimitAndOffset(page)
	keyCache := PREFIX_KEYCACHE_NEWS + "_" + bs.blogUtil.IntToString(limit) + "_" + bs.blogUtil.IntToString(offset)
	exist := bs.blogCache.Exists(keyCache)
	cache, err := bs.blogCache.Get(keyCache)
	var dbResult []domain.News
	if !exist || err != nil {
		dbResult, err = bs.blogRepo.GetNews(limit, offset)
		jsonByte := bs.blogUtil.ToJson(dbResult)
		jsonStr := bs.blogUtil.ByteToStr(jsonByte)
		bs.blogCache.Set(keyCache, string(jsonStr))
	} else {
		dbResult = bs.blogUtil.StringToArrayNews(cache)
	}
	return dbResult, err
}

func (bs BlogService) GetNewsDrafted(page int) ([]domain.News, error) {
	limit, offset := bs.blogUtil.GetLimitAndOffset(page)
	keyCache := PREFIX_KEYCACHE_NEWS_DRAFTED + "_" + bs.blogUtil.IntToString(limit) + "_" + bs.blogUtil.IntToString(offset)
	exist := bs.blogCache.Exists(keyCache)
	cache, err := bs.blogCache.Get(keyCache)
	var dbResult []domain.News
	if !exist || err != nil {
		dbResult, err = bs.blogRepo.GetNewsDrafted(limit, offset)
		jsonByte := bs.blogUtil.ToJson(dbResult)
		jsonStr := bs.blogUtil.ByteToStr(jsonByte)
		bs.blogCache.Set(keyCache, string(jsonStr))
	} else {
		dbResult = bs.blogUtil.StringToArrayNews(cache)
	}
	return dbResult, err
}

func (bs BlogService) GetNewsPublished(page int) ([]domain.News, error) {
	limit, offset := bs.blogUtil.GetLimitAndOffset(page)
	keyCache := PREFIX_KEYCACHE_NEWS_PUBLISHED + "_" + bs.blogUtil.IntToString(limit) + "_" + bs.blogUtil.IntToString(offset)
	exist := bs.blogCache.Exists(keyCache)
	cache, err := bs.blogCache.Get(keyCache)
	var dbResult []domain.News
	if !exist || err != nil {
		dbResult, err = bs.blogRepo.GetNewsPublished(limit, offset)
		jsonByte := bs.blogUtil.ToJson(dbResult)
		jsonStr := bs.blogUtil.ByteToStr(jsonByte)
		bs.blogCache.Set(keyCache, string(jsonStr))
	} else {
		dbResult = bs.blogUtil.StringToArrayNews(cache)
	}
	return dbResult, err
}

func (bs BlogService) GetNewsDeleted(page int) ([]domain.News, error) {
	limit, offset := bs.blogUtil.GetLimitAndOffset(page)
	keyCache := PREFIX_KEYCACHE_NEWS_DELETED + "_" + bs.blogUtil.IntToString(limit) + "_" + bs.blogUtil.IntToString(offset)
	exist := bs.blogCache.Exists(keyCache)
	cache, err := bs.blogCache.Get(keyCache)
	var dbResult []domain.News
	if !exist || err != nil {
		dbResult, err = bs.blogRepo.GetNewsDeleted(limit, offset)
		jsonByte := bs.blogUtil.ToJson(dbResult)
		jsonStr := bs.blogUtil.ByteToStr(jsonByte)
		bs.blogCache.Set(keyCache, string(jsonStr))
	} else {
		dbResult = bs.blogUtil.StringToArrayNews(cache)
	}
	return dbResult, err
}

func (bs BlogService) PostNews(new domain.News) error {
	err := bs.blogRepo.PostNews(new)
	if err == nil {
		var keys []string
		var err error
		prefixKey := ""
		if new.StatusContent == domain.DRAFTED {
			prefixKey = PREFIX_KEYCACHE_NEWS_DRAFTED
		} else if new.StatusContent == domain.PUBLISHED {
			prefixKey = PREFIX_KEYCACHE_NEWS_PUBLISHED
		} else if new.StatusContent == domain.DELETED {
			prefixKey = PREFIX_KEYCACHE_NEWS_DELETED
		}
		keyCache := prefixKey + "*"
		keys, err = bs.blogCache.Keys(keyCache)
		if err == nil {
			bs.blogEvent.PostNews(domain.News{
				Title:         "aku",
				Content:       "dia",
				StatusContent: domain.PUBLISHED,
			})
			for _, cache := range keys {
				go bs.blogCache.Del(cache)
			}
		}
	}
	return err
}

func (bs BlogService) GetSingleNews(id int64) (domain.News, error) {
	keyCache := PREFIX_KEYCACHE_SINGLE_NEWS + "_" + bs.blogUtil.Int64ToString(id)
	exist := bs.blogCache.Exists(keyCache)
	cache, err := bs.blogCache.Get(keyCache)
	var dbResult domain.News
	if !exist || err != nil {
		dbResult, err = bs.blogRepo.GetSingleNews(id)
		jsonByte := bs.blogUtil.ToJson(dbResult)
		jsonStr := bs.blogUtil.ByteToStr(jsonByte)
		bs.blogCache.Set(keyCache, string(jsonStr))
	} else {
		dbResult = bs.blogUtil.StringToNews(cache)
	}
	return dbResult, err
}

func (bs BlogService) PutSingleNews(new domain.News) error {
	err := bs.blogRepo.PutSingleNews(new)
	if err == nil {
		var keys []string
		var err error
		prefixKey := ""
		if new.StatusContent == domain.DRAFTED {
			prefixKey = PREFIX_KEYCACHE_NEWS_DRAFTED
		} else if new.StatusContent == domain.PUBLISHED {
			prefixKey = PREFIX_KEYCACHE_NEWS_PUBLISHED
		} else if new.StatusContent == domain.DELETED {
			prefixKey = PREFIX_KEYCACHE_NEWS_DELETED
		}
		keyCache := prefixKey + "*"
		keys, err = bs.blogCache.Keys(keyCache)
		if err == nil {
			for _, cache := range keys {
				go bs.blogCache.Del(cache)
			}
		}
	}
	return err
}

func (bs BlogService) DeleteSingleNews(id int64) error {
	dbResult, err := bs.blogRepo.GetSingleNews(id)
	if err != nil {
		return err
	}
	err = bs.blogRepo.DeleteSingleNews(id)
	if err == nil {
		keyCache := PREFIX_KEYCACHE_SINGLE_NEWS + "_" + bs.blogUtil.Int64ToString(id)
		exist := bs.blogCache.Exists(keyCache)
		if exist && err == nil {
			bs.blogCache.Del(keyCache)
		}
		var keys []string
		var err error
		prefixKey := ""
		if dbResult.StatusContent == domain.DRAFTED {
			prefixKey = PREFIX_KEYCACHE_NEWS_DRAFTED
		} else if dbResult.StatusContent == domain.PUBLISHED {
			prefixKey = PREFIX_KEYCACHE_NEWS_PUBLISHED
		} else if dbResult.StatusContent == domain.DELETED {
			prefixKey = PREFIX_KEYCACHE_NEWS_DELETED
		}
		keyCache = prefixKey + "*"
		keys, err = bs.blogCache.Keys(keyCache)
		if err == nil {
			for _, cache := range keys {
				go bs.blogCache.Del(cache)
			}
		}
	}
	return err
}

func (bs BlogService) PatchSingleNewsStatusContent(id int64, statusContent domain.StatusContent) error {
	err := bs.blogRepo.PatchSingleNewsStatusContent(id, statusContent)
	if err == nil {
		keyCache := PREFIX_KEYCACHE_SINGLE_NEWS + "_" + bs.blogUtil.Int64ToString(id)
		exist := bs.blogCache.Exists(keyCache)
		if exist {
			bs.blogCache.Del(keyCache)
		}
	}
	return err
}

func (bs BlogService) GetTopics(page int) ([]domain.Topics, error) {
	limit, offset := bs.blogUtil.GetLimitAndOffset(page)
	keyCache := PREFIX_KEYCACHE_TOPICS + "_" + bs.blogUtil.IntToString(limit) + "_" + bs.blogUtil.IntToString(offset)
	exist := bs.blogCache.Exists(keyCache)
	cache, err := bs.blogCache.Get(keyCache)
	var dbResult []domain.Topics
	if !exist || err != nil {
		dbResult, err = bs.blogRepo.GetTopics(limit, offset)
		jsonByte := bs.blogUtil.ToJson(dbResult)
		jsonStr := bs.blogUtil.ByteToStr(jsonByte)
		bs.blogCache.Set(keyCache, string(jsonStr))
	} else {
		dbResult = bs.blogUtil.StringToArrayTopics(cache)
	}
	return dbResult, err
}

func (bs BlogService) GetSingleTopics(id int64) (domain.Topics, error) {
	keyCache := PREFIX_KEYCACHE_SINGLE_TOPICS + "_" + bs.blogUtil.Int64ToString(id)
	exist := bs.blogCache.Exists(keyCache)
	cache, err := bs.blogCache.Get(keyCache)
	var dbResult domain.Topics
	if !exist || err != nil {
		dbResult, err = bs.blogRepo.GetSingleTopics(id)
		jsonByte := bs.blogUtil.ToJson(dbResult)
		jsonStr := bs.blogUtil.ByteToStr(jsonByte)
		bs.blogCache.Set(keyCache, string(jsonStr))
	} else {
		dbResult = bs.blogUtil.StringToTopics(cache)
	}
	return dbResult, err
}

func (bs BlogService) PostTopics(topic domain.Topics) error {
	err := bs.blogRepo.PostTopics(topic)
	if err == nil {
		var keys []string
		var err error
		keyCache := PREFIX_KEYCACHE_TOPICS + "*"
		keys, err = bs.blogCache.Keys(keyCache)
		if err == nil {
			for _, cache := range keys {
				go bs.blogCache.Del(cache)
			}
		}
	}
	return err
}

func (bs BlogService) PatchSingleTopicsCategoryName(idTopic int64, categoryName string) error {
	err := bs.blogRepo.PatchSingleTopicsCategoryName(idTopic, categoryName)
	if err == nil {
		keyCache := PREFIX_KEYCACHE_SINGLE_TOPICS + "_" + bs.blogUtil.Int64ToString(idTopic)
		exist := bs.blogCache.Exists(keyCache)
		if exist {
			bs.blogCache.Del(keyCache)
		}
	}
	return err
}

func (bs BlogService) GetSingleTopicsNews(idTopic int64, page int) ([]domain.News, error) {
	limit, offset := bs.blogUtil.GetLimitAndOffset(page)
	keyCache := PREFIX_KEYCACHE_SINGLE_TOPICS_NEWS + "_" + bs.blogUtil.IntToString(limit) + "_" + bs.blogUtil.IntToString(offset)
	exist := bs.blogCache.Exists(keyCache)
	cache, err := bs.blogCache.Get(keyCache)
	var dbResult []domain.News
	if !exist || err != nil {
		dbResult, err = bs.blogRepo.GetSingleTopicsNews(idTopic, limit, offset)
		jsonByte := bs.blogUtil.ToJson(dbResult)
		jsonStr := bs.blogUtil.ByteToStr(jsonByte)
		bs.blogCache.Set(keyCache, string(jsonStr))
	} else {
		dbResult = bs.blogUtil.StringToArrayNews(cache)
	}
	return dbResult, err
}

func (bs BlogService) PostRefNewsTopics(refNewsTopics domain.RefNewsTopics) error {
	err := bs.blogRepo.PostRefNewsTopics(refNewsTopics)
	if err == nil {
		var keys []string
		var err error
		keyCache := PREFIX_KEYCACHE_SINGLE_TOPICS_NEWS + "*"
		keys, err = bs.blogCache.Keys(keyCache)
		if err == nil {
			for _, cache := range keys {
				go bs.blogCache.Del(cache)
			}
		}
	}
	return err
}

func (bs BlogService) DeleteRefNewsTopics(refNewsTopics domain.RefNewsTopics) error {
	err := bs.blogRepo.DeleteRefNewsTopics(refNewsTopics)
	if err == nil {
		var keys []string
		var err error
		keyCache := PREFIX_KEYCACHE_SINGLE_TOPICS_NEWS + "*"
		keys, err = bs.blogCache.Keys(keyCache)
		if err == nil {
			for _, cache := range keys {
				go bs.blogCache.Del(cache)
			}
		}
	}
	return err
}
