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
