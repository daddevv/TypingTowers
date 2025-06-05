package game

import "fmt"

// SkillCategory is the high level grouping for a skill tree node.
type SkillCategory int

const (
	SkillOffense SkillCategory = iota
	SkillDefense
	SkillTyping
	SkillAutomation
	SkillUtility
)

// String returns a human readable label for the category.
func (c SkillCategory) String() string {
	switch c {
	case SkillOffense:
		return "Offense"
	case SkillDefense:
		return "Defense"
	case SkillTyping:
		return "Typing"
	case SkillAutomation:
		return "Automation"
	case SkillUtility:
		return "Utility"
	default:
		return "Unknown"
	}
}

// SkillNode represents a single unlockable skill in the global skill tree.
type SkillNode struct {
	ID       string
	Name     string
	Category SkillCategory
	Cost     int
	Effects  map[string]float64
	Prereqs  []string
}

// SkillTree holds all skill nodes keyed by ID.
type SkillTree struct {
	Nodes map[string]*SkillNode
	order []string
}

// NodesByCategory returns a slice of skill nodes belonging to the given category.
func (t *SkillTree) NodesByCategory(cat SkillCategory) []*SkillNode {
	var out []*SkillNode
	for _, id := range t.order {
		n := t.Nodes[id]
		if n.Category == cat {
			out = append(out, n)
		}
	}
	return out
}

// GetPrerequisites returns the prerequisite IDs for a given node.
func (t *SkillTree) GetPrerequisites(id string) []string {
	if n, ok := t.Nodes[id]; ok {
		return append([]string(nil), n.Prereqs...)
	}
	return nil
}

// UnlockOrder returns a topological order of node IDs based on prerequisites.
func (t *SkillTree) UnlockOrder() []string {
	return append([]string(nil), t.order...)
}

func (t *SkillTree) validate() error {
	for id, n := range t.Nodes {
		for _, p := range n.Prereqs {
			if _, ok := t.Nodes[p]; !ok {
				return fmt.Errorf("node %s missing prereq %s", id, p)
			}
		}
	}
	visited := map[string]bool{}
	stack := map[string]bool{}
	t.order = nil
	var visit func(string) error
	visit = func(id string) error {
		if stack[id] {
			return fmt.Errorf("cycle detected at %s", id)
		}
		if visited[id] {
			return nil
		}
		stack[id] = true
		for _, p := range t.Nodes[id].Prereqs {
			if err := visit(p); err != nil {
				return err
			}
		}
		stack[id] = false
		visited[id] = true
		t.order = append(t.order, id)
		return nil
	}
	for id := range t.Nodes {
		if err := visit(id); err != nil {
			return err
		}
	}
	return nil
}

// SampleSkillTree returns a small predefined skill tree used for tests and
// early prototyping.
func SampleSkillTree() (*SkillTree, error) {
	nodes := []SkillNode{
		{
			ID:       "sharp_arrows",
			Name:     "Sharp Arrows",
			Category: SkillOffense,
			Cost:     10,
			Effects:  map[string]float64{"damage_mult": 1.1},
		},
		{
			ID:       "rapid_fire",
			Name:     "Rapid Fire",
			Category: SkillOffense,
			Cost:     20,
			Effects:  map[string]float64{"fire_rate_mult": 0.9},
			Prereqs:  []string{"sharp_arrows"},
		},
		{
			ID:       "reinforced_walls",
			Name:     "Reinforced Walls",
			Category: SkillDefense,
			Cost:     10,
			Effects:  map[string]float64{"hp_add": 25},
		},
		{
			ID:       "touch_typing",
			Name:     "Touch Typing",
			Category: SkillTyping,
			Cost:     5,
			Effects:  map[string]float64{"wpm_bonus": 5},
		},
		{
			ID:       "auto_collect",
			Name:     "Auto Collect",
			Category: SkillAutomation,
			Cost:     15,
			Effects:  map[string]float64{"auto_collect": 1},
			Prereqs:  []string{"touch_typing"},
		},
		{
			ID:       "quick_commands",
			Name:     "Quick Commands",
			Category: SkillUtility,
			Cost:     5,
			Effects:  map[string]float64{"hotkeys": 1},
		},
	}
	tree := &SkillTree{Nodes: map[string]*SkillNode{}}
	for i := range nodes {
		n := nodes[i]
		tree.Nodes[n.ID] = &n
	}
	if err := tree.validate(); err != nil {
		return nil, err
	}
	return tree, nil
}
