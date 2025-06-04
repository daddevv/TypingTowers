package game

// TechNode represents a single unlockable technology in the game.
type TechNode struct {
	Name        string
	Letters     []rune
	Achievement string
}

// TechTree manages sequential technology unlocks.
type TechTree struct {
	nodes []TechNode
	stage int
}

// DefaultTechTree returns the default progression for letter unlocks.
func DefaultTechTree() *TechTree {
	nodes := []TechNode{
		{Name: "Home Row", Letters: []rune{'f', 'j'}, Achievement: "Unlock F & J"},
		{Name: "Index Extensions", Letters: []rune{'d', 'k'}, Achievement: "Unlock D & K"},
		{Name: "Middle Fingers", Letters: []rune{'s', 'l'}, Achievement: "Unlock S & L"},
		{Name: "Ring Finger", Letters: []rune{'a'}, Achievement: "Unlock A"},
		{Name: "Inner Index", Letters: []rune{'g', 'h'}, Achievement: "Unlock G & H"},
		{Name: "Top Row Pinky", Letters: []rune{'q', 'p'}, Achievement: "Unlock Q & P"},
		{Name: "Top Row Middle", Letters: []rune{'e', 'i'}, Achievement: "Unlock E & I"},
		{Name: "Top Row Index", Letters: []rune{'r', 'u'}, Achievement: "Unlock R & U"},
		{Name: "Top Row Outer", Letters: []rune{'t', 'y'}, Achievement: "Unlock T & Y"},
		{Name: "Top Row Ring", Letters: []rune{'w', 'o'}, Achievement: "Unlock W & O"},
		{Name: "Bottom Center", Letters: []rune{'c', 'm'}, Achievement: "Unlock C & M"},
		{Name: "Bottom Index", Letters: []rune{'v', 'n'}, Achievement: "Unlock V & N"},
		{Name: "Bottom Outer", Letters: []rune{'x', 'z'}, Achievement: "Unlock X & Z"},
	}
	return &TechTree{nodes: nodes, stage: 0}
}

// UnlockNext returns the letters from the next tech node and advances the stage.
func (t *TechTree) UnlockNext() (letters []rune, achievement string) {
	if t.stage >= len(t.nodes) {
		return nil, ""
	}
	node := t.nodes[t.stage]
	t.stage++
	return node.Letters, node.Achievement
}

// Completed returns true if all tech nodes have been unlocked.
func (t *TechTree) Completed() bool {
	return t.stage >= len(t.nodes)
}
