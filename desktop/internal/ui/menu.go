package ui

type BaseMenu struct {
	screen      *BaseScreen
	isSelected  bool
	activeOption int
	options     []string
}

func (m *BaseMenu) Selected() bool {
	return m.isSelected
}

func (m *BaseMenu) ActiveOption() int {
	return m.activeOption
}

func (m *BaseMenu) Options() []string {
	return m.options
}