package game

import "testing"

func TestQueueFIFO(t *testing.T) {
	q := NewQueueManager()
	q.Enqueue(Word{Text: "one", Source: "farmer", Family: "Gathering"})
	q.Enqueue(Word{Text: "two", Source: "barracks", Family: "Military"})

	if q.Len() != 2 {
		t.Fatalf("expected queue length 2 got %d", q.Len())
	}
	w, ok := q.Peek()
	if !ok || w.Text != "one" || w.Source != "farmer" || w.Family != "Gathering" {
		t.Fatalf("unexpected first word: %+v ok=%v", w, ok)
	}
}

func TestQueueDequeueValidation(t *testing.T) {
	q := NewQueueManager()
	q.Enqueue(Word{Text: "alpha", Source: "farmer", Family: "Gathering"})

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
	q := NewQueueManager()
	f := NewFarmer()
	b := NewBarracks()
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
