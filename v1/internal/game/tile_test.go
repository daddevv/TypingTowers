package game

import "testing"

func TestTileAtPosition(t *testing.T) {
    x, y := tileAtPosition(TileSize*5+1, TopMargin+TileSize*3+2)
    if x != 5 || y != 3 {
        t.Errorf("expected tile (5,3) got (%d,%d)", x, y)
    }
}

func TestTilePosition(t *testing.T) {
    px, py := tilePosition(2, 7)
    if px != TileSize*2 || py != TopMargin+TileSize*7 {
        t.Errorf("expected position (%d,%d) got (%d,%d)", TileSize*2, TopMargin+TileSize*7, px, py)
    }
}
