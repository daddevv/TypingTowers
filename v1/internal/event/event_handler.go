package event

import "sync"

// EventHandler routes events from handler channels through the EventBus.
// Handlers register their outbound channels with Register so events are
// published to the bus automatically.
type EventHandler struct {
	bus  *EventBus
	quit chan struct{}
	wg   sync.WaitGroup
}

// NewEventHandler creates a new EventHandler using the provided bus. If bus is
// nil a new EventBus is created.
func NewEventHandler(bus *EventBus) *EventHandler {
	if bus == nil {
		bus = NewEventBus()
	}
	return &EventHandler{bus: bus, quit: make(chan struct{})}
}

// Register begins forwarding all events received on ch to the bus using the
// given event type.
func (h *EventHandler) Register(eventType string, ch <-chan Event) {
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		for {
			select {
			case evt, ok := <-ch:
				if !ok {
					return
				}
				h.bus.Publish(eventType, evt)
			case <-h.quit:
				return
			}
		}
	}()
}

// Subscribe registers a subscriber channel for a specific event type.
func (h *EventHandler) Subscribe(eventType string, ch chan Event) {
	h.bus.Subscribe(eventType, ch)
}

// Unsubscribe removes a subscriber channel for a specific event type.
func (h *EventHandler) Unsubscribe(eventType string, ch chan Event) {
	h.bus.Unsubscribe(eventType, ch)
}

// Stop stops all forwarding goroutines and waits for them to finish.
func (h *EventHandler) Stop() {
	close(h.quit)
	h.wg.Wait()
}
