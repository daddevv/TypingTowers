package game

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

// TechNode represents a single unlockable technology in the game.
type TechNode struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Cost        int            `json:"cost"`
	Letters     []rune         `json:"-"`
	RawLetters  []string       `json:"letters"`
	Modifiers   TowerModifiers `json:"effects"`
	Prereqs     []string       `json:"prereqs"`
	Achievement string         `json:"achievement"`
}

// TechTree manages unlockable technologies loaded from a file.
type TechTree struct {
	nodes    []*TechNode
	index    map[string]*TechNode
	unlocked map[string]bool
	stage    int
}

// LoadTechTree reads tech nodes from the given path. The file is expected to be
// YAML (JSON syntax is valid). If loading fails an error is returned.
func LoadTechTree(path string) (*TechTree, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parseTechTree(data)
}

// parseTechTree unmarshals JSON/YAML data into a TechTree instance.
func parseTechTree(data []byte) (*TechTree, error) {
	var raw []*TechNode
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	t := &TechTree{
		nodes:    make([]*TechNode, 0, len(raw)),
		index:    make(map[string]*TechNode),
		unlocked: make(map[string]bool),
	}
	for _, n := range raw {
		if n.ID == "" {
			return nil, errors.New("tech node missing id")
		}
		for _, l := range n.RawLetters {
			if len(l) == 0 {
				continue
			}
			r := []rune(strings.ToLower(l))[0]
			n.Letters = append(n.Letters, r)
		}
		t.nodes = append(t.nodes, n)
		t.index[n.ID] = n
	}
	return t, nil
}

// DefaultTechTree returns an in-memory tech tree matching the built-in
// progression used by earlier versions of the game.
func DefaultTechTree() *TechTree {
	nodes := []*TechNode{
		{ID: "home_row", Name: "Home Row", Cost: 0, RawLetters: []string{"f", "j"}, Achievement: "Unlock F & J"},
		{ID: "index_ext", Name: "Index Extensions", Cost: 10, RawLetters: []string{"d", "k"}, Achievement: "Unlock D & K", Modifiers: TowerModifiers{RangeMult: 1.05}, Prereqs: []string{"home_row"}},
		{ID: "middle_fingers", Name: "Middle Fingers", Cost: 15, RawLetters: []string{"s", "l"}, Achievement: "Unlock S & L", Modifiers: TowerModifiers{DamageMult: 1.1}, Prereqs: []string{"index_ext"}},
		{ID: "ring_finger", Name: "Ring Finger", Cost: 20, RawLetters: []string{"a"}, Achievement: "Unlock A", Modifiers: TowerModifiers{AmmoAdd: 1}, Prereqs: []string{"middle_fingers"}},
	}
	data, _ := json.Marshal(nodes)
	tree, _ := parseTechTree(data)
	return tree
}

// UnlockNext sequentially unlocks the next node and returns its rewards.
func (t *TechTree) UnlockNext() (letters []rune, achievement string, mods TowerModifiers) {
	for t.stage < len(t.nodes) && t.unlocked[t.nodes[t.stage].ID] {
		t.stage++
	}
	if t.stage >= len(t.nodes) {
		return nil, "", TowerModifiers{}
	}
	node := t.nodes[t.stage]
	t.unlocked[node.ID] = true
	t.stage++
	return node.Letters, node.Achievement, node.Modifiers
}

// Purchase attempts to unlock the node with the given id if all prerequisites
// are met and returns the node details on success.
func (t *TechTree) Purchase(id string, gold int) (letters []rune, achievement string, mods TowerModifiers, cost int, ok bool) {
	node, ok := t.index[id]
	if !ok || t.unlocked[id] || gold < node.Cost {
		return nil, "", TowerModifiers{}, 0, false
	}
	for _, p := range node.Prereqs {
		if !t.unlocked[p] {
			return nil, "", TowerModifiers{}, 0, false
		}
	}
	t.unlocked[id] = true
	return node.Letters, node.Achievement, node.Modifiers, node.Cost, true
}

// Available returns all purchasable nodes filtered by the provided search term.
func (t *TechTree) Available(search string, gold int) []*TechNode {
	var out []*TechNode
	lower := strings.ToLower(search)
	for _, n := range t.nodes {
		if t.unlocked[n.ID] || gold < n.Cost {
			continue
		}
		ok := true
		for _, p := range n.Prereqs {
			if !t.unlocked[p] {
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

// Completed returns true if all tech nodes have been unlocked.
func (t *TechTree) Completed() bool {
	for _, n := range t.nodes {
		if !t.unlocked[n.ID] {
			return false
		}
	}
	return true
}
