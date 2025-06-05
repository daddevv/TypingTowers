package game

// Barracks spawns footmen when its word is typed correctly.
type Barracks struct {
	BaseEntity
	word     string
	progress int
	game     *Game
}

// NewBarracks creates a new barracks at the given position.
func NewBarracks(g *Game, x, y float64) *Barracks {
	w, h := ImgBarracks.Bounds().Dx(), ImgBarracks.Bounds().Dy()
	return &Barracks{
		BaseEntity: BaseEntity{
			pos:          Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgBarracks,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
			static:       true,
		},
		word: "foot",
		game: g,
	}
}

// Update processes typing input and spawns footmen.
func (b *Barracks) Update(dt float64) error {
	if b.game == nil {
		return nil
	}
	typed := b.game.input.TypedChars()
	for _, r := range typed {
		if b.progress < len(b.word) && r == rune(b.word[b.progress]) {
			b.progress++
			b.game.typing.Record(true)
			if b.progress >= len(b.word) {
				f := NewFootman(b.game, b.pos.X, b.pos.Y)
				b.game.units = append(b.game.units, f)
				b.progress = 0
			}
		} else if b.progress > 0 {
			b.game.typing.Record(false)
			b.progress = 0
		}
	}
	return nil
}
