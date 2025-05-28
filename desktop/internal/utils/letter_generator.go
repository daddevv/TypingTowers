package utils

import (
	"math/rand"
)

func GenerateRandomLetters(count int, possible []string) []string {
	letters := make([]string, count)
	for i := range letters {
		letters[i] = possible[rand.Intn(len(possible))]
	}
	return letters
}
