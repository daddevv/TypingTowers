//go:build test

package game

import "unicode"

// Step advances the game state by the provided delta time without relying on
// Ebiten's real-time loop. It processes the global typing queue and updates
// buildings, units, and the base. This helper is compiled only for tests.
func (g *Game) Step(dt float64) error {
	if g == nil {
		return nil
	}

	if g.queue != nil {
		g.queue.Update(dt)
		if w, ok := g.queue.Peek(); ok {
			if g.queueJam {
				if g.input.Backspace() {
					g.queueJam = false
					g.queueIndex = 0
				}
			} else {
				for _, r := range g.input.TypedChars() {
					expected := rune(w.Text[g.queueIndex])
					if unicode.ToLower(r) == unicode.ToLower(expected) {
						g.queueIndex++
						g.typing.Record(true)
						if g.queueIndex >= len(w.Text) {
							g.queueIndex = 0
							dq, _ := g.queue.TryDequeue(w.Text)
							switch dq.Source {
							case "Farmer":
								g.farmer.OnWordCompleted(dq.Text, &g.resources)
							case "Barracks":
								if unit := g.barracks.OnWordCompleted(dq.Text); unit != nil {
									g.military.AddUnit(unit)
								}
							}
						}
					} else {
						g.typing.Record(false)
						g.MistypeFeedback()
						g.queueJam = true
						break
					}
				}
			}
		}
	}

	if g.farmer != nil {
		g.farmer.Update(dt)
	}
	if g.lumberjack != nil {
		g.lumberjack.Update(dt)
	}
	if g.miner != nil {
		g.miner.Update(dt)
	}
	if g.barracks != nil {
		g.barracks.Update(dt)
	}
	if g.military != nil {
		g.military.Update(dt)
	}
	if g.base != nil {
		g.base.Update(dt)
	}

	g.input.Reset()
	return nil
}
