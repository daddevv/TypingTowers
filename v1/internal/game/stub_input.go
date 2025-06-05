package game

// stubInput implements InputHandler for deterministic test input.
type stubInput struct {
	typed     []rune
	backspace bool
}

func (s *stubInput) TypedChars() []rune { ch := s.typed; s.typed = nil; return ch }
func (s *stubInput) Update()            {}
func (s *stubInput) Reset()             { s.typed = nil; s.backspace = false }
func (s *stubInput) Backspace() bool    { v := s.backspace; s.backspace = false; return v }
func (s *stubInput) Space() bool        { return false }
func (s *stubInput) Quit() bool         { return false }
func (s *stubInput) Reload() bool       { return false }
func (s *stubInput) Enter() bool        { return false }
func (s *stubInput) Left() bool         { return false }
func (s *stubInput) Right() bool        { return false }
func (s *stubInput) Up() bool           { return false }
func (s *stubInput) Down() bool         { return false }
func (s *stubInput) Build() bool        { return false }
func (s *stubInput) Save() bool         { return false }
func (s *stubInput) Load() bool         { return false }
func (s *stubInput) SelectTower() bool  { return false }
func (s *stubInput) Command() bool      { return false }
func (s *stubInput) TechMenu() bool     { return false }
func (s *stubInput) SkillMenu() bool    { return false }
func (s *stubInput) StatsPanel() bool   { return false }
