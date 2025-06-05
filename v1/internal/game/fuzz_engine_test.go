package game

import (
	"encoding/json"
	"runtime/debug"
	"testing"
)

// FuzzGameRandomInput runs the game step with random input bytes to ensure
// robustness against unexpected sequences. It uses Go's built-in fuzzing
// to feed arbitrary byte slices as typed characters or backspace events.
func FuzzGameRandomInput(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		g := NewGame()
		inp := &stubInput{}
		g.input = inp

		captureState := func() string {
			s := struct {
				Wave     int `json:"wave"`
				BaseHP   int `json:"base_hp"`
				QueueLen int `json:"queue_len"`
				Mobs     int `json:"mobs"`
				Towers   int `json:"towers"`
			}{
				Wave:     g.currentWave,
				BaseHP:   g.base.Health(),
				QueueLen: g.queue.Len(),
				Mobs:     len(g.mobs),
				Towers:   len(g.towers),
			}
			b, _ := json.MarshalIndent(s, "", "  ")
			return string(b)
		}

		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("panic: %v\nstate: %s\ninput: %v\n%s", r, captureState(), data, debug.Stack())
			}
		}()

		for i, b := range data {
			if b%50 == 0 {
				inp.backspace = true
			} else {
				inp.typed = []rune{rune('a' + b%26)}
			}
			if err := g.Step(0.05); err != nil {
				t.Fatalf("step error: %v\nstate: %s\ninput index %d: %v", err, captureState(), i, data[:i+1])
			}
			if i > 1000 {
				break // limit runtime for extreme inputs
			}
		}
	})
}
