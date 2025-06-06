package econ

// LetterStage defines letters unlocked at each stage and their King's Point cost.
type LetterStage struct {
	Letters []rune
	Cost    int
}

// LetterUnlockStages is the ordered progression for unlocking letters.
var LetterUnlockStages = []LetterStage{
	{Letters: []rune{'f', 'j'}, Cost: 0},
	{Letters: []rune{'d', 'k'}, Cost: 20},
	{Letters: []rune{'s', 'l'}, Cost: 40},
	{Letters: []rune{'a'}, Cost: 60},
	{Letters: []rune{'g', 'h'}, Cost: 90},
	{Letters: []rune{'q', 'p'}, Cost: 120},
	{Letters: []rune{'e', 'i'}, Cost: 150},
	{Letters: []rune{'r', 'u'}, Cost: 180},
	{Letters: []rune{'t', 'y'}, Cost: 210},
	{Letters: []rune{'w', 'o'}, Cost: 240},
	{Letters: []rune{'c', 'm'}, Cost: 270},
	{Letters: []rune{'v', 'n'}, Cost: 310},
	{Letters: []rune{'x', 'z'}, Cost: 350},
}

// LetterStageLetters returns the letters unlocked at the given stage.
func LetterStageLetters(stage int) []rune {
	if stage < 0 || stage >= len(LetterUnlockStages) {
		return nil
	}
	return LetterUnlockStages[stage].Letters
}

// LetterStageCost returns the King's Point cost for the given stage.
func LetterStageCost(stage int) int {
	if stage < 0 || stage >= len(LetterUnlockStages) {
		return -1
	}
	return LetterUnlockStages[stage].Cost
}
