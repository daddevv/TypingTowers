package game

import "testing"

func TestFootmanKillsGrunt(t *testing.T) {
	b := NewBase(0, 0, 10)
	g := &Game{mobs: []Enemy{}, units: []*Footman{}, input: NewInput(), typing: NewTypingStats()}
	g.base = b
	grunt := NewOrcGrunt(50, 0, b)
	g.mobs = []Enemy{grunt}
	foot := NewFootman(g, 0, 0)
	g.units = []*Footman{foot}

	steps := int(8.0/0.016) + 1
	for i := 0; i < steps && grunt.Alive(); i++ {
		foot.Update(0.016)
		grunt.Update(0.016)
		if !grunt.Alive() {
			break
		}
	}

	if grunt.Alive() {
		t.Errorf("footman did not kill grunt within 8 seconds")
	}
}
