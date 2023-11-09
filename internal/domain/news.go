package domain

import "time"

type StatusContent string

const (
	DRAFTED   StatusContent = "drafted"
	PUBLISHED StatusContent = "published"
	DELETED   StatusContent = "deleted"
)

type News struct {
	Id            int64         `json:"id,omitempty"`
	Title         string        `json:"title"`
	Content       string        `json:"content,omitempty"`
	StatusContent StatusContent `json:"status_content"`
	CreatedAt     time.Time     `json:"created_at,omitempty"`
	UpdatedAt     time.Time     `json:"updated_at,omitempty"`
}
