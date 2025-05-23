package game

type Menu struct {
	Options  []string
	Selected int
}

func NewMenu() *Menu {
	return &Menu{
		Options:  []string{"Start Game", "Options", "Quit"},
		Selected: 0,
	}
}

func (m *Menu) Update(input *InputState) {
	if input.Up {
		m.Selected--
		if m.Selected < 0 {
			m.Selected = len(m.Options) - 1
		}
	}
	if input.Down {
		m.Selected++
		if m.Selected >= len(m.Options) {
			m.Selected = 0
		}
	}
	if input.Enter {
		m.Select()
	}
}

func (m *Menu) Select() {
	switch m.Selected {
	case 0:
		// Start Game
	case 1:
		// Open Options
	case 2:
		// Quit
	}
}
