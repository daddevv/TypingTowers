//go:build test

package event

import "testing"

func TestEventBusPublishSubscribe(t *testing.T) {
	bus := NewEventBus()
	ch := make(chan Event, 1)
	bus.Subscribe("ui", ch)

	evt := UIEvent{Type: "notification", Payload: "hello"}
	bus.Publish("ui", evt)

	select {
	case rec := <-ch:
		uevt, ok := rec.(UIEvent)
		if !ok {
			t.Fatalf("expected UIEvent got %T", rec)
		}
		if uevt.Type != "notification" || uevt.Payload != "hello" {
			t.Fatalf("unexpected event data: %+v", uevt)
		}
	default:
		t.Fatalf("no event received")
	}
}

func TestEventBusDropsWhenChannelFull(t *testing.T) {
	bus := NewEventBus()
	ch := make(chan Event, 1)
	bus.Subscribe("ui", ch)

	bus.Publish("ui", UIEvent{Type: "n1"})
	bus.Publish("ui", UIEvent{Type: "n2"})

	<-ch
	select {
	case <-ch:
		t.Fatalf("expected second event to be dropped")
	default:
	}
}
