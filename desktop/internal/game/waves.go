package game

import "td/internal/entity"

type WaveConfig struct {
	PossibleLetters []string `json:"possible_letters"` // Possible letters for the wave
	MinLength     int      `json:"min_length"` // Minimum length of words in the wave
	MaxLength     int      `json:"max_length"` // Maximum length of words in the wave
	PossibleEnemies []entity.Mob `json:"possible_enemies"` // Possible enemies for the wave
	EnemyCount	 	int          `json:"enemy_count"` // Number of enemies to defeat before the wave ends
}