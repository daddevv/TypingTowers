package entity

import (
	"math/rand"
	"td/internal/utils"
)

// MobSpawner handles the periodic spawning of mobs with configurable timing and randomization.
type MobSpawner struct {
	// Core timing
	MinSpawnInterval float64 // Minimum seconds between spawns
	MaxSpawnInterval float64 // Maximum seconds between spawns
	NextSpawnTime    float64 // Time until next spawn (in seconds)
	ElapsedTime      float64 // Total elapsed time since spawner creation

	// Spawn configuration
	LetterPool   LetterPool // LetterPool for dynamic letter sets
	WordPool     []string   // Pool of specific words that can be spawned
	RandomWeight float64    // Weight for spawning random letter mobs (0.0-1.0)
	WordWeight   float64    // Weight for spawning word mobs (0.0-1.0)
	
	// Letter count for random mobs
	MinLetterCount int // Minimum letters for random mobs
	MaxLetterCount int // Maximum letters for random mobs

	// Mob factories for extensibility
	MobFactories []func([]string) Mob
}

// NewMobSpawner creates a new MobSpawner with default settings and a LetterPool.
func NewMobSpawner(letterPool LetterPool) *MobSpawner {
	spawner := &MobSpawner{
		MinSpawnInterval: 3.0,
		MaxSpawnInterval: 5.0,
		LetterPool:       letterPool,
		RandomWeight:     1.0,
		WordWeight:       0.0,
		MinLetterCount:   2,
		MaxLetterCount:   4,
		WordPool:         []string{},
		MobFactories:     []func([]string) Mob{func(letters []string) Mob { return NewBeachballMobWithLetters(letters) }},
	}
	spawner.resetSpawnTimer()
	return spawner
}

// Update advances the spawner's timer and returns a new mob if it's time to spawn one.
// deltaTime should be provided in seconds.
func (ms *MobSpawner) Update(deltaTime float64, score int) Mob {
	ms.ElapsedTime += deltaTime
	ms.NextSpawnTime -= deltaTime
	ms.LetterPool.Update(score)
	if ms.NextSpawnTime <= 0 {
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
func (ms *MobSpawner) spawnMob() Mob {
	choice := rand.Float64()
	possibleLetters := ms.LetterPool.GetPossibleLetters()
	if choice < ms.RandomWeight {
		letterCount := ms.MinLetterCount + rand.Intn(ms.MaxLetterCount-ms.MinLetterCount+1)
		letters := utils.GenerateRandomLetters(letterCount, possibleLetters)
		// Use first registered mob factory
		return ms.MobFactories[0](letters)
	} else if len(ms.WordPool) > 0 {
		word := ms.WordPool[rand.Intn(len(ms.WordPool))]
		letters := make([]string, len(word))
		for i, char := range word {
			letters[i] = string(char)
		}
		return ms.MobFactories[0](letters)
	} else {
		letterCount := ms.MinLetterCount + rand.Intn(ms.MaxLetterCount-ms.MinLetterCount+1)
		letters := utils.GenerateRandomLetters(letterCount, possibleLetters)
		return ms.MobFactories[0](letters)
	}
}

// RegisterMobFactory allows adding new mob types for spawning.
func (ms *MobSpawner) RegisterMobFactory(factory func([]string) Mob) {
	ms.MobFactories = append(ms.MobFactories, factory)
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
func (ms *MobSpawner) ForceSpawn() Mob {
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
