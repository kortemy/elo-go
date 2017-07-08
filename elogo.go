package elogo

import (
	"math"
)

const (
	// K is the default K-Factor
	K = 32
	// D is the default deviation
	D = 400
)

// Elo calculates Elo rating changes based on the configured factors.
type Elo struct {
	K int
	D int
}

// Outcome is a match result data for a single player.
type Outcome struct {
	Delta int
	Rating int
}

// NewElo instantiates the Elo object with default factors.
// Default K-Factor is 32
// Default deviation is 400 
func NewElo() *Elo {
	return &Elo{K, D}
}

// NewEloWithFactors instantiates the Elo object with custom factor values.
func NewEloWithFactors(k, d int) *Elo {
	return &Elo{k, d}
}

// ExpectedScore gives the expected chance that the first player wins
func (e *Elo) ExpectedScore(ratingA, ratingB int) float64 {
	return e.ExpectedScoreWithFactors(ratingA, ratingB, e.D)
}

// ExpectedScoreWithFactors overrides default factors and gives the expected chance that the first player wins
func (e *Elo) ExpectedScoreWithFactors(ratingA, ratingB, d int) float64 {
	return 1 / (1 + math.Pow(10, float64(ratingB - ratingA) / float64(d)))
}

// RatingDelta gives the ratings change for the first player for the given score
func (e *Elo) RatingDelta(ratingA, ratingB int, score float64) int {
	return e.RatingDeltaWithFactors(ratingA, ratingB, score, e.K, e.D)
}

// RatingDeltaWithFactors overrides default factors and gives the ratings change for the first player for the given score
func (e *Elo) RatingDeltaWithFactors(ratingA, ratingB int, score float64, k, d int ) int {
	return int(float64(k) * (score - e.ExpectedScoreWithFactors(ratingA, ratingB, d)))
}

// Rating gives the new rating for the first player for the given score
func (e *Elo) Rating(ratingA, ratingB int, score float64) int {
	return e.RatingWithFactors(ratingA, ratingB, score, e.K, e.D)
}
// RatingWithFactors overrides default factors and gives the new rating for the first player for the given score
func (e *Elo) RatingWithFactors(ratingA, ratingB int, score float64, k, d int ) int {
	return ratingA + e.RatingDeltaWithFactors(ratingA, ratingB, score, k, d)
}

// Outcome gives an Outcome object for each player for the given score
func (e *Elo) Outcome(ratingA, ratingB int, score float64) (Outcome, Outcome) {
	return e.OutcomeWithFactors(ratingA, ratingB, score, e.K, e.D)
}

// OutcomeWithFactors overrides default factors and gives an Outcome object for each player for the given score
func (e *Elo) OutcomeWithFactors(ratingA, ratingB int, score float64, k, d int ) (Outcome, Outcome) {
	delta := e.RatingDeltaWithFactors(ratingA, ratingB, score, k, d)
	return Outcome{ delta, ratingA + delta }, Outcome{ -delta, ratingB - delta }
}