package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

type stubInputSkill struct{ stubInputOverlay }

func TestDrawSkillTreeOverlay(t *testing.T) {
	g := NewGame()
	g.input = &stubInputSkill{}
	g.skillMenuOpen = true
	hud := NewHUD(g)
	img := ebiten.NewImage(1920, 1080)
	hud.drawSkillTreeOverlay(img)
}
