package elo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestNewRating(t *testing.T) {
// 	expectedRating := 1516
// 	var testRating float64 = 1500
// 	testResult := 1

// 	newRating := NewRating(testRating, testResult)
// 	assert.Equal(t, newRating, expectedRating, "Expect new rating to be 1516.")
// }

func TestNewRatingWin(t *testing.T) {
	var newRating float64 = 1516
	var oldRating float64 = 1500
	matchRes := 1.0
	winProb := 0.500

	result := NewRating(oldRating, winProb, matchRes)
	assert.Equal(t, newRating, result,
		"Expect new rating to be 1516.")
}

func TestNewRatingLoss(t *testing.T) {
	var newRating float64 = 1484
	var oldRating float64 = 1500
	matchRes := 0.0
	winProb := 0.500

	result := NewRating(oldRating, winProb, matchRes)
	assert.Equal(t, newRating, result,
		"Expect new rating to be 1516.")
}

func TestNewRatingDrawFavorite(t *testing.T) {
	var newRating float64 = 1492
	var oldRating float64 = 1500
	matchRes := 0.5
	winProb := 0.750

	result := NewRating(oldRating, winProb, matchRes)
	assert.Equal(t, newRating, result,
		"Expect new rating to be 1516.")
}

func TestNewRatingDrawUnderdog(t *testing.T) {
	var newRating float64 = 1510
	var oldRating float64 = 1500
	matchRes := 0.5
	winProb := 0.200

	result := NewRating(oldRating, winProb, matchRes)
	assert.Equal(t, newRating, result,
		"Expect new rating to be 1516.")
}

func TestWinProbabilityEven(t *testing.T) {
	expectedProbability := 0.500
	var testRating float64 = 1500
	var testOpponent float64 = 1500

	result := WinProbability(testRating, testOpponent)
	assert.Equal(t, expectedProbability, result,
		"Expect win probablity to be 50.0%.")
}

func TestWinProbabilityFavorite(t *testing.T) {
	expectedProbability := 0.909
	var testRating float64 = 1700
	var testOpponent float64 = 1300

	result := WinProbability(testRating, testOpponent)
	assert.Equal(t, expectedProbability, result,
		"Expect win probablity to be 90.9%.")
}

func TestWinProbabilityUnderdog(t *testing.T) {
	expectedProbability := 0.267
	var testRating float64 = 1250
	var testOpponent float64 = 1425

	result := WinProbability(testRating, testOpponent)
	assert.Equal(t, expectedProbability, result,
		"Expect win probablity to be 9.1%.")
}
