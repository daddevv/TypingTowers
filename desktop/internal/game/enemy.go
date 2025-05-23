package game

type Enemy struct {
	Health int
	Speed  int
}

func NewEnemy(health, speed int) *Enemy {
	return &Enemy{
		Health: health,
		Speed:  speed,
	}
}
