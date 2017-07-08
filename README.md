# Elo for Go

Elo ranking system for Go (Golang) programming language.

[Elo ranking system](https://en.wikipedia.org/wiki/Elo_rating_system) is a system for ranking players in a competitive environment, by calculating expected outcome in a match between two ranked players and their new ranks based on the outcome of a match. Can be used in head-to-head mathes, as well as team based matches.

## Factors

There are two factors that affect how Elo works - *K-factor* and *deviation*.

**Deviation** controls how much the difference in ratings affects the _expected_ outcome. The higher the number, the higher expected chance that a better rated player wins. Default deviation is 400. 

**K-factor** controls how much points are gained/lost after the outcome of the match is known. The higher the factor, players rating is more affected. This is especially useful if you want to control point fluctuation based on number of games played. Players with less games played should have higher _k-factor_, because their early ranking doesn't necessarily replicate their skill level. Default k-factor is 24.

## Documentation

[Full documentation on GoDoc.](https://godoc.org/github.com/kortemy/elo-go)

## Sample usage

```
rankA := 1500
rankB := 1600
elo := NewElo() // or NewEloWithFactors(k, d) for custom factors

// Expected chance that A defeats B
// use ExpectedScoreWithFactors(rankA, rankB, deviation) to use custom factor for this method
elo.ExpectedScore(rankA, rankB) // 0.3599350001971149

// Results for A in the outcome of A defeats B
score := 1 // Use 1 in case A wins, 0 in case B wins, 0.5 in case of a draw
elo.RatingDelta(rankA, rankB, score) // 20
elo.Rating(rankA, rankB, score) // 1520
outcomeA, outcomeB := elo.Outcome(rankA, rankB, score)
outcomeA.Delta // 20
outcomeA.Rating // 1520
outcomeB.Delta // -20
outcomeB.Rating // 1580
```

All `Elo` methods have `WithFactor` variants, where you can override factors for that specific method call:
```
kFactor = 40
deviation := 800

elo.ExpectedScore(rankA, rankB, deviation)
elo.RatingDelta(rankA, rankB, score, kFactor, deviation)
elo.Rating(rankA, rankB, score, kFactor, deviation)
elo.Outcome(rankA, rankB, score, kFactor, deviation)
```
