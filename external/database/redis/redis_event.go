package redis

import (
	"dot-golang/internal/domain"
	"fmt"
	"time"
)

type EventBus struct {
	subscribers map[string][]chan<- domain.Event
}

type EventService struct {
	eventBus *EventBus
}

func PostNewsEventHandler(eventChan <-chan domain.Event) {
	for event := range eventChan {
		new, ok := event.Data.(domain.News)
		if !ok {
			continue
		}
		fmt.Println("New has been created with ID:", new.Id)
	}
}

// NewEventBus creates a new instance of the event bus
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan<- domain.Event),
	}
}

// Subscribe adds a new subscriber for a given event type
func (eb *EventBus) Subscribe(eventType string, subscriber chan<- domain.Event) {
	eb.subscribers[eventType] = append(eb.subscribers[eventType], subscriber)
}

// Publish sends an event to all subscribers of a given event type
func (eb *EventBus) Publish(event domain.Event) {
	subscribers := eb.subscribers[event.Type]
	for _, subscriber := range subscribers {
		subscriber <- event
	}
}

// NewEventPostNewsService creates a new instance of the user registration service
func NewEventPostNewsService(eventBus *EventBus) *EventService {
	return &EventService{
		eventBus: eventBus,
	}
}

// RegisterUser registers a new user and publishes a user registered event
func (urs *EventService) PostNews(new domain.News) {
	event := domain.Event{
		Type:      "UserRegistered",
		Timestamp: time.Now(),
		Data:      new,
	}

	// Publish the event
	urs.eventBus.Publish(event)
}
