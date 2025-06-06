//go:build test

package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/event"
)

type stubHandler struct{ called bool }

func (s *stubHandler) Update(dt float64) { s.called = true }

func TestHandlersUpdateCalled(t *testing.T) {
	g := NewGame()
	g.input = &mockInput{}
	g.phase = core.PhasePlaying

	e := &stubHandler{}
	u := &stubHandler{}
	te := &stubHandler{}
	to := &stubHandler{}
	ph := &stubHandler{}
	c := &stubHandler{}
	s := &stubHandler{}

	g.EntityHandler = e
	g.UIHandler = u
	g.TechHandler = te
	g.TowerHandler = to
	g.PhaseHandler = ph
	// g.ContentHandler = c

	if err := g.Update(); err != nil {
		t.Fatal(err)
	}

	if !e.called || !u.called || !te.called || !to.called || !ph.called || !c.called || !s.called {
		t.Fatalf("expected all handlers to be updated")
	}
}

func TestApplyNextTechEmitsUIEvent(t *testing.T) {
	g := NewGame()
	g.UIEvents = make(chan event.Event, 1)

	g.applyNextTech()

	select {
	case evt := <-g.UIEvents:
		if uevt, ok := evt.(event.UIEvent); !ok || uevt.Type != "notification" {
			t.Fatalf("unexpected event: %#v", evt)
		}
	default:
		t.Fatalf("no UI event received")
	}
}
