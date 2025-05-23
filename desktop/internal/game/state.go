package game

type State struct {
	CurrentScreen string
	Score         int
}

func NewState() *State {
	return &State{
		CurrentScreen: "main-menu",
		Score:         0,
	}
}

func (s *State) Reset() {
	s.CurrentScreen = "main-menu"
	s.Score = 0
}
