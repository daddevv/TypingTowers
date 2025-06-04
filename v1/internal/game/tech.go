package game

// TechNode represents a single unlockable technology in the game.
type TechNode struct {
	Name        string
	Letters     []rune
	Modifiers   TowerModifiers
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
		{Name: "Index Extensions", Letters: []rune{'d', 'k'}, Achievement: "Unlock D & K", Modifiers: TowerModifiers{RangeMult: 1.05}},
		{Name: "Middle Fingers", Letters: []rune{'s', 'l'}, Achievement: "Unlock S & L", Modifiers: TowerModifiers{DamageMult: 1.1}},
		{Name: "Ring Finger", Letters: []rune{'a'}, Achievement: "Unlock A", Modifiers: TowerModifiers{AmmoAdd: 1}},
		{Name: "Inner Index", Letters: []rune{'g', 'h'}, Achievement: "Unlock G & H", Modifiers: TowerModifiers{FireRateMult: 0.95}},
		{Name: "Top Row Pinky", Letters: []rune{'q', 'p'}, Achievement: "Unlock Q & P", Modifiers: TowerModifiers{DamageMult: 1.1}},
		{Name: "Top Row Middle", Letters: []rune{'e', 'i'}, Achievement: "Unlock E & I", Modifiers: TowerModifiers{RangeMult: 1.05}},
		{Name: "Top Row Index", Letters: []rune{'r', 'u'}, Achievement: "Unlock R & U", Modifiers: TowerModifiers{AmmoAdd: 1}},
		{Name: "Top Row Outer", Letters: []rune{'t', 'y'}, Achievement: "Unlock T & Y", Modifiers: TowerModifiers{FireRateMult: 0.95}},
		{Name: "Top Row Ring", Letters: []rune{'w', 'o'}, Achievement: "Unlock W & O", Modifiers: TowerModifiers{DamageMult: 1.1}},
		{Name: "Bottom Center", Letters: []rune{'c', 'm'}, Achievement: "Unlock C & M", Modifiers: TowerModifiers{RangeMult: 1.05}},
		{Name: "Bottom Index", Letters: []rune{'v', 'n'}, Achievement: "Unlock V & N", Modifiers: TowerModifiers{AmmoAdd: 1}},
		{Name: "Bottom Outer", Letters: []rune{'x', 'z'}, Achievement: "Unlock X & Z", Modifiers: TowerModifiers{FireRateMult: 0.95}},
	}
	return &TechTree{nodes: nodes, stage: 0}
}

// UnlockNext returns the letters from the next tech node and advances the stage.
func (t *TechTree) UnlockNext() (letters []rune, achievement string, mods TowerModifiers) {
	if t.stage >= len(t.nodes) {
		return nil, "", TowerModifiers{}
	}
	node := t.nodes[t.stage]
	t.stage++
	return node.Letters, node.Achievement, node.Modifiers
}

// Completed returns true if all tech nodes have been unlocked.
func (t *TechTree) Completed() bool {
	return t.stage >= len(t.nodes)
}
