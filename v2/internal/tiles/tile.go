package tiles

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	ID       int           // Unique identifier for the tile
	Name     string        // Name of the tile
	Image    *ebiten.Image // Path to the tile image
	Walkable bool          // Indicates if the tile can be walked on
	// Add other fields as necessary, such as tile properties, animations, etc.
}

// NewTile creates a new Tile with the specified properties.
func NewTile(id int, name string, image *ebiten.Image, walkable bool) *Tile {
	return &Tile{
		ID:       id,
		Name:     name,
		Image:    image,
		Walkable: walkable,
	}
}
