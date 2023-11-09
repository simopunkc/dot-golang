package abstraction

import "dot-golang/internal/domain"

type BlogUtil interface {
	GetLimitAndOffset(page int) (int, int)
	StringToInt(raw string) (int, error)
	IntToString(raw int) string
	Int64ToString(raw int64) string
	StringToInt64(raw string) (int64, error)
	ToJson(raw interface{}) []byte
	ByteToStr(raw []byte) string
	StringToNews(jsonStr string) domain.News
	StringToTopics(jsonStr string) domain.Topics
	StringToArrayNews(jsonStr string) []domain.News
	StringToArrayTopics(jsonStr string) []domain.Topics
}
