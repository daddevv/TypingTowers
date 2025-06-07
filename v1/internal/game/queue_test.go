package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/building"
	"github.com/daddevv/type-defense/internal/building/gatherer"
	"github.com/daddevv/type-defense/internal/core"
)

func TestQueueFIFO(t *testing.T) {
	q := core.NewQueueManager()
	q.Enqueue(core.Word{Text: "one", Source: "farmer", Family: "Gathering"})
	q.Enqueue(core.Word{Text: "two", Source: "barracks", Family: "Military"})

	if q.Len() != 2 {
		t.Fatalf("expected queue length 2 got %d", q.Len())
	}
	w, ok := q.Peek()
	if !ok || w.Text != "one" || w.Source != "farmer" || w.Family != "Gathering" {
		t.Fatalf("unexpected first word: %+v ok=%v", w, ok)
	}
}

func TestQueueDequeueValidation(t *testing.T) {
	q := core.NewQueueManager()
	q.Enqueue(core.Word{Text: "alpha", Source: "farmer", Family: "Gathering"})

	if _, ok := q.TryDequeue("beta"); ok {
		t.Fatalf("dequeue should fail for wrong input")
	}
	if q.Len() != 1 {
		t.Fatalf("queue length changed on failed dequeue")
	}
	w, ok := q.TryDequeue("alpha")
	if !ok || w.Text != "alpha" || w.Family != "Gathering" {
		t.Fatalf("dequeue failed for correct input")
	}
	if q.Len() != 0 {
		t.Fatalf("expected empty queue after dequeue")
	}
}

func TestQueueEnqueueFromBuildings(t *testing.T) {
	q := core.NewQueueManager()
	f := gatherer.NewFarmer()
	b := building.NewBarracks()
	f.SetQueue(q)
	b.SetQueue(q)

	f.SetInterval(0.1)
	f.SetCooldown(0.1)
	b.SetInterval(0.1)
	b.SetCooldown(0.1)

	f.Update(0.11)
	b.Update(0.11)

	if q.Len() != 2 {
		t.Fatalf("expected 2 words in queue got %d", q.Len())
	}
}

func TestQueueBackPressureDamage(t *testing.T) {
	q := core.NewQueueManager()
	base := building.NewBase(0, 0, 5)
	for i := 0; i < 6; i++ {
		q.Enqueue(core.Word{Text: "w"})
	}
	q.Update(1.0, base) // Pass base to apply damage
	if base.Health() != 4 {
		t.Fatalf("expected base health 4 got %d", base.Health())
	}
}

func TestColorize(t *testing.T) {
	w := core.Word{Text: "foo", Family: "Gathering"}
	got := w.Colorize()
	expected := "\033[32mfoo\033[0m"
	if got != expected {
		t.Fatalf("Colorize mismatch: got %q want %q", got, expected)
	}
}
