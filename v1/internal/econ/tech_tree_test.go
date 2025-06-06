package econ

import (
	"os"
	"testing"
)

func TestLoadTechTree(t *testing.T) {
	yaml := `nodes:
  - id: home_row
    name: Home Row
    type: UnlockLetter
    cost: 0
    letters: [f, j]
    prereqs: []
  - id: index_ext
    name: Index Extensions
    type: UnlockLetter
    cost: 20
    letters: [d, k]
    prereqs: [home_row]
  - id: middle_fingers
    name: Middle Fingers
    type: UnlockLetter
    cost: 30
    letters: [s, l]
    prereqs: [index_ext]
  - id: ring_finger
    name: Ring Finger
    type: UnlockLetter
    cost: 40
    letters: [a, ;]
    prereqs: [middle_fingers]
`
	tmp, err := os.CreateTemp("", "techtree*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.Write([]byte(yaml)); err != nil {
		t.Fatal(err)
	}
	tmp.Close()
	tree, err := LoadTechTree(tmp.Name())
	if err != nil {
		t.Fatalf("load tech tree: %v", err)
	}
	if len(tree.Nodes) != 4 {
		t.Fatalf("expected 4 nodes got %d", len(tree.Nodes))
	}
	node, ok := tree.Nodes["index_ext"]
	if !ok {
		t.Fatalf("missing node index_ext")
	}
	if node.Cost != 20 {
		t.Fatalf("expected cost 20 got %d", node.Cost)
	}
}

func TestValidateCycle(t *testing.T) {
	yaml := `nodes:
  - id: a
    name: A
    type: UnlockLetter
    cost: 0
    prereqs: [b]
  - id: b
    name: B
    type: UnlockLetter
    cost: 0
    prereqs: [a]
`
	tmp, err := os.CreateTemp("", "cycle*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.Write([]byte(yaml)); err != nil {
		t.Fatal(err)
	}
	tmp.Close()
	if _, err := LoadTechTree(tmp.Name()); err == nil {
		t.Fatalf("expected cycle error")
	}
}

func TestUnlockOrder(t *testing.T) {
	yaml := `nodes:
  - id: home_row
    name: Home Row
    type: UnlockLetter
    cost: 0
    letters: [f, j]
    prereqs: []
  - id: index_ext
    name: Index Extensions
    type: UnlockLetter
    cost: 20
    letters: [d, k]
    prereqs: [home_row]
  - id: middle_fingers
    name: Middle Fingers
    type: UnlockLetter
    cost: 30
    letters: [s, l]
    prereqs: [index_ext]
  - id: ring_finger
    name: Ring Finger
    type: UnlockLetter
    cost: 40
    letters: [a, ;]
    prereqs: [middle_fingers]
`
	tmp, err := os.CreateTemp("", "techtree*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	if _, err := tmp.Write([]byte(yaml)); err != nil {
		t.Fatal(err)
	}
	tmp.Close()
	tree, err := LoadTechTree(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}
	order := tree.UnlockOrder()
	expected := []string{"home_row", "index_ext", "middle_fingers", "ring_finger"}
	if len(order) != len(expected) {
		t.Fatalf("unexpected order length %v", order)
	}
	for i, id := range expected {
		if order[i] != id {
			t.Fatalf("expected %s at %d got %s", id, i, order[i])
		}
	}
}
