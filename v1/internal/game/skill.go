package game

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

// SkillNode represents a skill unlock in the global skill tree.
type SkillNode struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Category    string         `json:"category"`
	ReqWPM      float64        `json:"req_wpm"`
	ReqAccuracy float64        `json:"req_accuracy"`
	Modifiers   TowerModifiers `json:"effects"`
	Prereqs     []string       `json:"prereqs"`
}

// SkillTree manages unlockable skills.
type SkillTree struct {
	nodes    []*SkillNode
	index    map[string]*SkillNode
	unlocked map[string]bool
}

// LoadSkillTree reads a skill tree file from disk.
func LoadSkillTree(path string) (*SkillTree, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parseSkillTree(data)
}

// parseSkillTree unmarshals JSON/YAML data into a SkillTree instance.
func parseSkillTree(data []byte) (*SkillTree, error) {
	var raw []*SkillNode
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	st := &SkillTree{
		nodes:    make([]*SkillNode, 0, len(raw)),
		index:    make(map[string]*SkillNode),
		unlocked: make(map[string]bool),
	}
	for _, n := range raw {
		if n.ID == "" {
			return nil, errors.New("skill node missing id")
		}
		st.nodes = append(st.nodes, n)
		st.index[n.ID] = n
	}
	return st, nil
}

// DefaultSkillTree returns the built in skill tree used for tests.
func DefaultSkillTree() *SkillTree {
	nodes := []*SkillNode{
		{ID: "offense_basic", Name: "Sharpened Tips", Category: "offense", Modifiers: TowerModifiers{DamageMult: 1.05}},
		{ID: "typing_focus", Name: "Focus Training", Category: "typing", ReqWPM: 30, ReqAccuracy: 0.9, Modifiers: TowerModifiers{FireRateMult: 0.95}, Prereqs: []string{"offense_basic"}},
		{ID: "automation_reload", Name: "Auto-Reload", Category: "automation", ReqWPM: 40, ReqAccuracy: 0.95, Modifiers: TowerModifiers{AmmoAdd: 2}, Prereqs: []string{"typing_focus"}},
	}
	data, _ := json.Marshal(nodes)
	st, _ := parseSkillTree(data)
	return st
}

// Unlock attempts to unlock a skill node by id if requirements are met.
func (st *SkillTree) Unlock(id string, stats TypingStats) (TowerModifiers, bool) {
	node, ok := st.index[id]
	if !ok || st.unlocked[id] {
		return TowerModifiers{}, false
	}
	if stats.WPM() < node.ReqWPM || stats.Accuracy() < node.ReqAccuracy {
		return TowerModifiers{}, false
	}
	for _, p := range node.Prereqs {
		if !st.unlocked[p] {
			return TowerModifiers{}, false
		}
	}
	st.unlocked[id] = true
	return node.Modifiers, true
}

// Available returns all skills that can currently be unlocked.
func (st *SkillTree) Available(search string, stats TypingStats) []*SkillNode {
	var out []*SkillNode
	lower := strings.ToLower(search)
	for _, n := range st.nodes {
		if st.unlocked[n.ID] {
			continue
		}
		if stats.WPM() < n.ReqWPM || stats.Accuracy() < n.ReqAccuracy {
			continue
		}
		ok := true
		for _, p := range n.Prereqs {
			if !st.unlocked[p] {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		if lower != "" && !strings.Contains(strings.ToLower(n.Name), lower) {
			continue
		}
		out = append(out, n)
	}
	return out
}

// Unlocked returns a slice of IDs for all unlocked skills.
func (st *SkillTree) Unlocked() []string {
	var ids []string
	for id, ok := range st.unlocked {
		if ok {
			ids = append(ids, id)
		}
	}
	return ids
}

// SetUnlocked marks the provided skill ids as unlocked. Unknown ids are ignored.
func (st *SkillTree) SetUnlocked(ids []string) {
	for _, id := range ids {
		if _, ok := st.index[id]; ok {
			st.unlocked[id] = true
		}
	}
}
