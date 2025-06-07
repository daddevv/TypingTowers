//go:build test

package event

import "testing"

func TestEventHandlerRoutesEvents(t *testing.T) {
	bus := NewEventBus()
	h := NewEventHandler(bus)
	pub := make(chan Event, 1)
	sub := make(chan Event, 1)

	h.Register("entity", pub)
	h.Subscribe("entity", sub)

	pub <- EntityEvent{Type: "spawn", Payload: "orc"}

	select {
	case e := <-sub:
		ev, ok := e.(EntityEvent)
		if !ok || ev.Type != "spawn" || ev.Payload != "orc" {
			t.Fatalf("unexpected event: %#v", e)
		}
	default:
		t.Fatalf("no event routed")
	}

	h.Stop()
}
