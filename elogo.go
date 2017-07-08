package elogo

import (
	"math"
)

const (
	K = 32
	D = 400
)

type Elo struct {
	K int
	D int
}

type Outcome struct {
	Delta int
	Rating int
}

func NewElo() *Elo {
	return &Elo{K, D}
}

func NewEloWithFactors(k, d int) *Elo {
	return &Elo{k, d}
}

func (e *Elo) ExpectedScore(ratingA, ratingB int) float64 {
	return e.ExpectedScoreWithFactors(ratingA, ratingB, e.D)
}

func (e *Elo) ExpectedScoreWithFactors(ratingA, ratingB, d int) float64 {
	return 1 / (1 + math.Pow(10, float64(ratingB - ratingA) / float64(d)))
}

func (e *Elo) RatingDelta(ratingA, ratingB int, outcome float64) int {
	return e.RatingDeltaWithFactors(ratingA, ratingB, outcome, e.K, e.D)
}

func (e *Elo) RatingDeltaWithFactors(ratingA, ratingB int, outcome float64, k, d int ) int {
	return int(float64(k) * (outcome - e.ExpectedScoreWithFactors(ratingA, ratingB, d)))
}

func (e *Elo) Rating(ratingA, ratingB int, outcome float64) int {
	return e.RatingWithFactors(ratingA, ratingB, outcome, e.K, e.D)
}

func (e *Elo) RatingWithFactors(ratingA, ratingB int, outcome float64, k, d int ) int {
	return ratingA + e.RatingDeltaWithFactors(ratingA, ratingB, outcome, k, d)
}

func (e *Elo) Outcome(ratingA, ratingB int, outcome float64) (Outcome, Outcome) {
	return e.OutcomeWithFactors(ratingA, ratingB, outcome, e.K, e.D)
}

func (e *Elo) OutcomeWithFactors(ratingA, ratingB int, outcome float64, k, d int ) (Outcome, Outcome) {
	delta := e.RatingDeltaWithFactors(ratingA, ratingB, outcome, k, d)
	return Outcome{ delta, ratingA + delta }, Outcome{ -delta, ratingB - delta }
}