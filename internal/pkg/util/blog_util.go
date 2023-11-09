package util

import (
	"bytes"
	"dot-golang/internal/domain"
	"encoding/json"
	"strconv"
)

type BlogUtil struct {
}

func NewBlogUtil() *BlogUtil {
	return &BlogUtil{}
}

func (bu BlogUtil) GetLimitAndOffset(page int) (int, int) {
	limit := 10
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	return limit, offset
}

func (bu BlogUtil) StringToInt(raw string) (int, error) {
	result, err := strconv.Atoi(raw)
	return result, err
}

func (bu BlogUtil) IntToString(raw int) string {
	result := strconv.Itoa(raw)
	return result
}

func (bu BlogUtil) Int64ToString(raw int64) string {
	result := strconv.FormatInt(raw, 10)
	return result
}

func (bu BlogUtil) StringToInt64(raw string) (int64, error) {
	result, err := strconv.ParseInt(raw, 10, 64)
	return result, err
}

func (bu BlogUtil) ToJson(raw interface{}) []byte {
	result, err := json.Marshal(raw)
	if err != nil {
		return []byte{}
	}
	return result
}

func (bu BlogUtil) ByteToStr(raw []byte) string {
	return bytes.NewBuffer(raw).String()
}

func (bu BlogUtil) StringToNews(jsonStr string) domain.News {
	var news domain.News
	err := json.Unmarshal([]byte(jsonStr), &news)
	if err != nil {
		return domain.News{}
	}
	return news
}

func (bu BlogUtil) StringToTopics(jsonStr string) domain.Topics {
	var topics domain.Topics
	err := json.Unmarshal([]byte(jsonStr), &topics)
	if err != nil {
		return domain.Topics{}
	}
	return topics
}

func (bu BlogUtil) StringToArrayNews(jsonStr string) []domain.News {
	var news []domain.News
	err := json.Unmarshal([]byte(jsonStr), &news)
	if err != nil {
		return []domain.News{}
	}
	return news
}

func (bu BlogUtil) StringToArrayTopics(jsonStr string) []domain.Topics {
	var topics []domain.Topics
	err := json.Unmarshal([]byte(jsonStr), &topics)
	if err != nil {
		return []domain.Topics{}
	}
	return topics
}
