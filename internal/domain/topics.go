package domain

import "time"

type Topics struct {
	Id           int64     `json:"id,omitempty"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
