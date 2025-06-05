package game

import (
	"os"
	"testing"
)

func TestLoadTechTree(t *testing.T) {
	tmp, err := os.CreateTemp("", "tree*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	data := `[{"id":"a","name":"A","cost":1,"letters":["f"],"effects":{},"prereqs":[]}]`
	if _, err := tmp.Write([]byte(data)); err != nil {
		t.Fatal(err)
	}
	tmp.Close()
	tree, err := LoadTechTree(tmp.Name())
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if len(tree.nodes) != 1 {
		t.Fatalf("expected 1 node")
	}
}
