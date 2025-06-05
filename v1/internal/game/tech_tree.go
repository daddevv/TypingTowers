package game

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// NodeEffects describes modifiers granted by a tech node.
type NodeEffects struct {
	Letters    []string `yaml:"letters"`
	RangeMult  float64  `yaml:"range_mult"`
	DamageMult float64  `yaml:"damage_mult"`
	AmmoAdd    int      `yaml:"ammo_add"`
}

// YAMLTechNode represents a node in the tech tree YAML.
type YAMLTechNode struct {
	ID      string      `yaml:"id"`
	Name    string      `yaml:"name"`
	Type    string      `yaml:"type"`
	Cost    int         `yaml:"cost"`
	Effects NodeEffects `yaml:"effects"`
	Prereqs []string    `yaml:"prereqs"`
}

type yamlTechFile struct {
	Nodes []YAMLTechNode `yaml:"nodes"`
}

// YAMLTechTree is an in-memory graph of technology nodes loaded from YAML.
type YAMLTechTree struct {
	Nodes map[string]*YAMLTechNode
	order []string
}

// LoadTechTree parses a YAML file into a YAMLTechTree and validates it.
func LoadTechTree(path string) (*YAMLTechTree, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var f yamlTechFile
	if err := yaml.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	tree := &YAMLTechTree{Nodes: map[string]*YAMLTechNode{}}
	for i := range f.Nodes {
		n := f.Nodes[i]
		tree.Nodes[n.ID] = &n
	}
	if err := tree.validate(); err != nil {
		return nil, err
	}
	return tree, nil
}

// GetPrerequisites returns the prerequisite IDs for a node.
func (t *YAMLTechTree) GetPrerequisites(id string) []string {
	if n, ok := t.Nodes[id]; ok {
		return append([]string(nil), n.Prereqs...)
	}
	return nil
}

// UnlockOrder returns a topological order of node IDs.
func (t *YAMLTechTree) UnlockOrder() []string {
	return append([]string(nil), t.order...)
}

// validate checks for missing prereqs and cycles and builds UnlockOrder.
func (t *YAMLTechTree) validate() error {
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
