package abstraction

type BlogEvent interface {
	EventPublisher(eventType string, data interface{})
}
