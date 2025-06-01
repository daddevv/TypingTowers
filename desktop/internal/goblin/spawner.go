package goblin

import (
	"td/internal/enemy"
	"td/internal/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type GoblinSpawner struct {
	Sprite    *ebiten.Image // Sprite for the enemy spawner
	Position  *math.Vec2    // Position of the spawner
	SpawnRate int           // Rate at which enemies are spawned (in ticks)
	SpawnTick int           // Current tick for spawning enemies
}

func NewGoblinSpawner(sprite *ebiten.Image, x, y, spawnRate int) *GoblinSpawner {
	return &GoblinSpawner{
		Sprite:    sprite,
		Position:  math.NewVec2(float64(x), float64(y)),
		SpawnRate: spawnRate,
		SpawnTick: 0, // Initialize spawn tick
	}
}

func (es *GoblinSpawner) Update(enemyList []*enemy.Mob) []*enemy.Mob {
	// Increment the spawn tick
	es.SpawnTick++

	// Check if it's time to spawn an enemy
	if es.SpawnTick >= es.SpawnRate {
		es.SpawnTick = 0     // Reset the spawn tick
		es.NewMob(enemyList) // Call the method to spawn an enemy
	}

	return enemyList
}

func (es *GoblinSpawner) NewMob(enemyList []*enemy.Mob) []*enemy.Mob {
	// Create a new enemy mob at the spawner's position
	newEnemy := enemy.NewMob("Goblin", 100, int(es.Position.X), int(es.Position.Y), 0, 0, 100, 100)
	newEnemy.ActionRange = 100   // Set action range for the enemy
	newEnemy.AttackDamage = 10   // Set attack damage for the enemy
	newEnemy.AttackCooldown = 60 // Set attack cooldown in ticks

	// Add the new enemy to the list
	enemyList = append(enemyList, newEnemy)

	return enemyList
}
