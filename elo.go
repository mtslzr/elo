package elo

import (
	"math"
)

// Default k Factor
const kF = 32

// NewRating takes an old rating and returns a new rating, based on result.
func NewRating(rating float64, prob float64, result float64) float64 {
	new := kF * (result - prob)
	return math.Round(float64(rating) + new)
}

// WinProbability takes two teams' ratings and returns chance of winning.
func WinProbability(team float64, oppo float64) float64 {
	diff := -(team - oppo) / 400
	prob := 1 / (math.Pow(10, diff) + 1)
	return math.Round(prob*1000) / 1000
}
