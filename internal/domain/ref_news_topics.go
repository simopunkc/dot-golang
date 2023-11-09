package domain

type RefNewsTopics struct {
	Id       int64 `json:"id,omitempty"`
	NewsId   int64 `json:"news_id"`
	TopicsId int64 `json:"topics_id"`
}
