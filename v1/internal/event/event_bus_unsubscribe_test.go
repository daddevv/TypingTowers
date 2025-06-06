//go:build test

package event

import "testing"

func TestEventBusUnsubscribe(t *testing.T) {
	bus := NewEventBus()
	ch := make(chan Event, 1)
	bus.Subscribe("ui", ch)
	bus.Unsubscribe("ui", ch)
	bus.Publish("ui", UIEvent{Type: "notice"})
	select {
	case <-ch:
		t.Fatalf("received event after unsubscribe")
	default:
	}
}
