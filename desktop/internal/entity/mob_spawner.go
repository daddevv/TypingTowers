package entity

import (
	"math/rand"
)

// MobSpawner handles the periodic spawning of mobs with configurable timing and randomization.
type MobSpawner struct {
	// Core timing
	MinSpawnInterval float64 // Minimum seconds between spawns
	MaxSpawnInterval float64 // Maximum seconds between spawns
	NextSpawnTime    float64 // Time until next spawn (in seconds)
	ElapsedTime      float64 // Total elapsed time since spawner creation

	// Spawn configuration
	PossibleLetters []string // Available letters for random mobs
	WordPool        []string // Pool of specific words that can be spawned
	RandomWeight    float64  // Weight for spawning random letter mobs (0.0-1.0)
	WordWeight      float64  // Weight for spawning word mobs (0.0-1.0)
	
	// Letter count for random mobs
	MinLetterCount int // Minimum letters for random mobs
	MaxLetterCount int // Maximum letters for random mobs
}

// NewMobSpawner creates a new MobSpawner with default settings.
func NewMobSpawner(possibleLetters []string) *MobSpawner {
	spawner := &MobSpawner{
		MinSpawnInterval: 3.0,  // 3 seconds minimum (slower start)
		MaxSpawnInterval: 5.0,  // 5 seconds maximum (slower start)
		PossibleLetters:  possibleLetters,
		RandomWeight:     1.0,  // 100% chance for random letters
		WordWeight:       0.0,  // 0% chance for words
		MinLetterCount:   2,    // Minimum 2 letters
		MaxLetterCount:   4,    // Maximum 4 letters (reduced from 5)
		WordPool:        []string{}, // Empty word pool
	}
	
	// Initialize the next spawn time
	spawner.resetSpawnTimer()
	
	return spawner
}

// Update advances the spawner's timer and returns a new mob if it's time to spawn one.
// deltaTime should be provided in seconds.
func (ms *MobSpawner) Update(deltaTime float64) Entity {
	ms.ElapsedTime += deltaTime
	ms.NextSpawnTime -= deltaTime
	
	if ms.NextSpawnTime <= 0 {
		// Time to spawn a new mob
		ms.resetSpawnTimer()
		return ms.spawnMob()
	}
	
	return nil
}

// resetSpawnTimer sets the next spawn time to a random value within the configured range.
func (ms *MobSpawner) resetSpawnTimer() {
	interval := ms.MinSpawnInterval + rand.Float64()*(ms.MaxSpawnInterval-ms.MinSpawnInterval)
	ms.NextSpawnTime = interval
}

// spawnMob creates and returns a new mob based on the spawner's configuration.
func (ms *MobSpawner) spawnMob() Entity {
	// Determine whether to spawn a random letter mob or a word mob
	choice := rand.Float64()
	
	if choice < ms.RandomWeight {
		// Spawn random letter mob
		letterCount := ms.MinLetterCount + rand.Intn(ms.MaxLetterCount-ms.MinLetterCount+1)
		return NewBeachballMob(letterCount, ms.PossibleLetters)
	} else if len(ms.WordPool) > 0 {
		// Spawn word mob
		word := ms.WordPool[rand.Intn(len(ms.WordPool))]
		return NewBeachballMobWithWord(word)
	} else {
		// Fallback to random letters if no words available
		letterCount := ms.MinLetterCount + rand.Intn(ms.MaxLetterCount-ms.MinLetterCount+1)
		return NewBeachballMob(letterCount, ms.PossibleLetters)
	}
}

// SetSpawnInterval configures the spawn timing range.
func (ms *MobSpawner) SetSpawnInterval(minSeconds, maxSeconds float64) {
	ms.MinSpawnInterval = minSeconds
	ms.MaxSpawnInterval = maxSeconds
}

// SetLetterCount configures the letter count range for random mobs.
func (ms *MobSpawner) SetLetterCount(min, max int) {
	ms.MinLetterCount = min
	ms.MaxLetterCount = max
}

// SetSpawnWeights configures the probability of spawning random vs word mobs.
// randomWeight + wordWeight should ideally sum to 1.0 for proper probability distribution.
func (ms *MobSpawner) SetSpawnWeights(randomWeight, wordWeight float64) {
	ms.RandomWeight = randomWeight
	ms.WordWeight = wordWeight
}

// AddWordsToPool adds new words to the spawner's word pool.
func (ms *MobSpawner) AddWordsToPool(words []string) {
	ms.WordPool = append(ms.WordPool, words...)
}

// ClearWordPool removes all words from the spawner's word pool.
func (ms *MobSpawner) ClearWordPool() {
	ms.WordPool = ms.WordPool[:0]
}

// SetWordPool replaces the current word pool with the provided words.
func (ms *MobSpawner) SetWordPool(words []string) {
	ms.WordPool = make([]string, len(words))
	copy(ms.WordPool, words)
}

// GetTimeUntilNextSpawn returns the time in seconds until the next mob spawn.
func (ms *MobSpawner) GetTimeUntilNextSpawn() float64 {
	return ms.NextSpawnTime
}

// ForceSpawn immediately spawns a mob and resets the spawn timer.
func (ms *MobSpawner) ForceSpawn() Entity {
	ms.resetSpawnTimer()
	return ms.spawnMob()
}

// SpeedUpOverTime gradually decreases spawn intervals based on score.
// Call this method when mobs are defeated to increase difficulty.
func (ms *MobSpawner) SpeedUpOverTime(score int) {
	// Reduce spawn time by 0.1 seconds for every 10 mobs defeated
	// Minimum spawn time is 1.0 second
	speedBonus := float64(score) * 0.01 // 0.01 seconds per mob defeated
	
	ms.MinSpawnInterval = max(1.0, 3.0-speedBonus)
	ms.MaxSpawnInterval = max(1.5, 5.0-speedBonus)
}

// max returns the larger of two float64 values
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
