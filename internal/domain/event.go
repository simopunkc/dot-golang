package domain

import "time"

type Event struct {
	Type      string
	Timestamp time.Time
	Data      interface{}
}
