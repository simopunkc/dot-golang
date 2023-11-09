package redis

import (
	"dot-golang/internal/abstraction"
	"dot-golang/internal/constant"
	"dot-golang/internal/domain"
	"time"
)

type EventBus struct {
	subscribers map[string][]chan<- domain.Event
}

type EventService struct {
	eventBus *EventBus
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan<- domain.Event),
	}
}

func (eb *EventBus) Subscribe(eventType string, subscriber chan<- domain.Event) {
	eb.subscribers[eventType] = append(eb.subscribers[eventType], subscriber)
}

func (eb *EventBus) Publish(event domain.Event) {
	subscribers := eb.subscribers[event.Type]
	for _, subscriber := range subscribers {
		subscriber <- event
	}
}

func NewEventService(eventBus *EventBus) *EventService {
	return &EventService{
		eventBus: eventBus,
	}
}

func (urs *EventService) EventPublisher(eventType string, data interface{}) {
	event := domain.Event{
		Type:      eventType,
		Timestamp: time.Now(),
		Data:      data,
	}
	urs.eventBus.Publish(event)
}

func EventConsumer(blogCache abstraction.BlogCache, blogUtil abstraction.BlogUtil, eventChan <-chan domain.Event) {
	for event := range eventChan {
		if event.Type == constant.EVENT_DELETE_SINGLE_NEWS {
			new, ok := event.Data.(domain.News)
			if !ok {
				continue
			}
			keyCache := constant.PREFIX_KEYCACHE_SINGLE_NEWS + "_" + blogUtil.Int64ToString(new.Id)
			blogCache.Del(keyCache)
		}
	}
}
