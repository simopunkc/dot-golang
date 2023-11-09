package repository

import (
	"dot-golang/internal/constant"
	"dot-golang/internal/domain"
	"time"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db}
}

func (br BlogRepository) GetNews(limit int, offset int) ([]domain.News, error) {
	var news []domain.News
	err := br.db.Table(constant.TABLE_NEWS).Limit(limit).Offset(offset).Find(&news).Error
	return news, err
}

func (br BlogRepository) GetNewsDrafted(limit int, offset int) ([]domain.News, error) {
	var news []domain.News
	err := br.db.Table(constant.TABLE_NEWS).Where(constant.WHERE_COLUMN_STATUS_CONTENT, domain.DRAFTED).Limit(limit).Offset(offset).Find(&news).Error
	return news, err
}

func (br BlogRepository) GetNewsPublished(limit int, offset int) ([]domain.News, error) {
	var news []domain.News
	err := br.db.Table(constant.TABLE_NEWS).Where(constant.WHERE_COLUMN_STATUS_CONTENT, domain.PUBLISHED).Limit(limit).Offset(offset).Find(&news).Error
	return news, err
}

func (br BlogRepository) GetNewsDeleted(limit int, offset int) ([]domain.News, error) {
	var news []domain.News
	err := br.db.Table(constant.TABLE_NEWS).Where(constant.WHERE_COLUMN_STATUS_CONTENT, domain.DELETED).Limit(limit).Offset(offset).Find(&news).Error
	return news, err
}

func (br BlogRepository) PostNews(new domain.News) error {
	err := br.db.Save(&new).Error
	return err
}

func (br BlogRepository) GetSingleNews(id int64) (domain.News, error) {
	var new domain.News
	err := br.db.Table(constant.TABLE_NEWS).Where(constant.WHERE_COLUMN_ID, id).First(&new).Error
	return new, err
}

func (br BlogRepository) PutSingleNews(new domain.News) error {
	new.UpdatedAt = time.Now()
	err := br.db.Save(&new).Error
	return err
}

func (br BlogRepository) DeleteSingleNews(id int64) error {
	var new domain.News
	var err error
	err = br.db.Table(constant.TABLE_NEWS).Where(constant.WHERE_COLUMN_ID, id).First(&new).Error
	if err != nil {
		return err
	}

	new.StatusContent = domain.DELETED
	new.UpdatedAt = time.Now()
	err = br.db.Save(&new).Error
	return err
}

func (br BlogRepository) PatchSingleNewsStatusContent(id int64, statusContent domain.StatusContent) error {
	var new domain.News
	var err error
	err = br.db.Table(constant.TABLE_NEWS).Where(constant.WHERE_COLUMN_ID, id).First(&new).Error
	if err != nil {
		return err
	}

	new.StatusContent = statusContent
	new.UpdatedAt = time.Now()
	err = br.db.Save(&new).Error
	return err
}

func (br BlogRepository) GetTopics(limit int, offset int) ([]domain.Topics, error) {
	var topics []domain.Topics
	err := br.db.Table(constant.TABLE_TOPICS).Limit(limit).Offset(offset).Find(&topics).Error
	return topics, err
}

func (br BlogRepository) GetSingleTopics(id int64) (domain.Topics, error) {
	var topic domain.Topics
	err := br.db.Table(constant.TABLE_TOPICS).Where(constant.WHERE_COLUMN_ID, id).First(&topic).Error
	return topic, err
}

func (br BlogRepository) PostTopics(topic domain.Topics) error {
	err := br.db.Save(&topic).Error
	return err
}

func (br BlogRepository) PatchSingleTopicsCategoryName(idTopic int64, categoryName string) error {
	var topic domain.Topics
	var err error
	err = br.db.Table(constant.TABLE_TOPICS).Where(constant.WHERE_COLUMN_ID, idTopic).First(&topic).Error
	if err != nil {
		return err
	}

	topic.CategoryName = categoryName
	topic.UpdatedAt = time.Now()
	err = br.db.Save(&topic).Error
	return err
}

func (br BlogRepository) GetSingleTopicsNews(idTopic int64, limit int, offset int) ([]domain.News, error) {
	var news []domain.News
	err := br.db.Table(constant.TABLE_REF_NEWS_TOPICS).Where(constant.WHERE_COLUMN_TOPICS_ID, idTopic).Limit(limit).Offset(offset).Find(&news).Error
	return news, err
}

func (br BlogRepository) PostRefNewsTopics(refNewsTopics domain.RefNewsTopics) error {
	tx := br.db.Begin()
	err := tx.Create(&refNewsTopics).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}

func (br BlogRepository) DeleteRefNewsTopics(refNewsTopics domain.RefNewsTopics) error {
	var topic domain.Topics
	var err error
	err = br.db.Table(constant.TABLE_REF_NEWS_TOPICS).Where(constant.WHERE_COLUMN_TOPICS_ID, refNewsTopics.TopicsId).Where(constant.WHERE_COLUMN_NEWS_ID, refNewsTopics.NewsId).First(&topic).Error
	if err != nil {
		return err
	}

	tx := br.db.Begin()
	err = tx.Table(constant.TABLE_REF_NEWS_TOPICS).Delete(topic).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return err
}
