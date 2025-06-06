package game

// mockInput implements InputHandler for deterministic test input.
type mockInput struct {
	typed     []rune
	backspace bool
}

func (s *mockInput) TypedChars() []rune { ch := s.typed; s.typed = nil; return ch }
func (s *mockInput) Update()            {}
func (s *mockInput) Reset()             { s.typed = nil; s.backspace = false }
func (s *mockInput) Backspace() bool    { v := s.backspace; s.backspace = false; return v }
func (s *mockInput) Space() bool        { return false }
func (s *mockInput) Quit() bool         { return false }
func (s *mockInput) Reload() bool       { return false }
func (s *mockInput) Enter() bool        { return false }
func (s *mockInput) Left() bool         { return false }
func (s *mockInput) Right() bool        { return false }
func (s *mockInput) Up() bool           { return false }
func (s *mockInput) Down() bool         { return false }
func (s *mockInput) Build() bool        { return false }
func (s *mockInput) Save() bool         { return false }
func (s *mockInput) Load() bool         { return false }
func (s *mockInput) SelectTower() bool  { return false }
func (s *mockInput) Command() bool      { return false }
func (s *mockInput) TechMenu() bool     { return false }
func (s *mockInput) SkillMenu() bool    { return false }
func (s *mockInput) StatsPanel() bool   { return false }
