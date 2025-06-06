package game

import "testing"

func TestSampleSkillTree(t *testing.T) {
	tree, err := SampleSkillTree()
	if err != nil {
		t.Fatalf("sample skill tree: %v", err)
	}
	if len(tree.Nodes) != 6 {
		t.Fatalf("expected 6 nodes got %d", len(tree.Nodes))
	}
	order := tree.UnlockOrder()
	if len(order) != 6 {
		t.Fatalf("unexpected unlock order length %d", len(order))
	}
	// ensure prerequisite ordering
	for i, id := range order {
		if id == "rapid_fire" && i == 0 {
			t.Fatalf("rapid_fire should come after sharp_arrows")
		}
		if id == "auto_collect" {
			// touch_typing must appear before auto_collect
			seen := false
			for _, prev := range order[:i] {
				if prev == "touch_typing" {
					seen = true
					break
				}
			}
			if !seen {
				t.Fatalf("auto_collect missing prereq before it")
			}
		}
	}
}

func TestSkillTreeCycleDetect(t *testing.T) {
	tree := &SkillTree{Nodes: map[string]*SkillNode{
		"a": {ID: "a", Prereqs: []string{"b"}},
		"b": {ID: "b", Prereqs: []string{"a"}},
	}}
	if err := tree.validate(); err == nil {
		t.Fatalf("expected cycle error")
	}
}

func TestSkillUnlockFlow(t *testing.T) {
	tree, err := SampleSkillTree()
	if err != nil {
		t.Fatalf("sample tree: %v", err)
	}
	pool := &ResourcePool{}
	pool.AddKingsPoints(50)

	// cannot unlock rapid_fire before sharp_arrows
	if tree.Unlock("rapid_fire", pool) {
		t.Fatalf("rapid_fire unlocked without prereq")
	}

	if !tree.Unlock("sharp_arrows", pool) {
		t.Fatalf("failed to unlock sharp_arrows")
	}
	if !tree.unlocked["sharp_arrows"] {
		t.Fatalf("sharp_arrows not marked unlocked")
	}

	if pool.KingsAmount() != 40 {
		t.Fatalf("expected 40 KP remaining got %d", pool.KingsAmount())
	}

	if !tree.Unlock("rapid_fire", pool) {
		t.Fatalf("failed to unlock rapid_fire after prereq")
	}
	if pool.KingsAmount() != 20 {
		t.Fatalf("expected 20 KP remaining got %d", pool.KingsAmount())
	}
	if !tree.unlocked["rapid_fire"] {
		t.Fatalf("rapid_fire not marked unlocked")
	}
}
