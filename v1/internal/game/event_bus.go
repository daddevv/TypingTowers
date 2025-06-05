package game

import "sync"

// EventBus provides a simple pub/sub mechanism for events.
type EventBus struct {
	subscribers map[string][]chan Event
	mu          sync.RWMutex
}

// NewEventBus creates a new EventBus.
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan Event),
	}
}

// Subscribe registers a channel to receive events of a given type.
func (eb *EventBus) Subscribe(eventType string, ch chan Event) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subscribers[eventType] = append(eb.subscribers[eventType], ch)
}

// Publish sends an event to all subscribers of its type.
func (eb *EventBus) Publish(eventType string, evt Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	for _, ch := range eb.subscribers[eventType] {
		select {
		case ch <- evt:
		default:
			// Drop event if channel is full
		}
	}
}
