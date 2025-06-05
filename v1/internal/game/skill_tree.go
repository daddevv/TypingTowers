package game

// SkillCategory is the high level grouping for a skill tree node.
type SkillCategory int

const (
	SkillOffense SkillCategory = iota
	SkillDefense
	SkillTyping
	SkillAutomation
	SkillUtility
)

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
}
