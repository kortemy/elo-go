package elogo

import (
	"testing"
)

const (
	A = 1500
	B = 1600
)

func TestElo(t *testing.T) {
	elo := NewElo()

	expected := elo.ExpectedScore(A, B)
	if expected != 0.3599350001971149 {
		t.Fatalf("Expected chance %v, but got %v", 0.3599350001971149, expected)
	}

	delta1 := elo.RatingDelta(A, B, 1)
	if delta1 != 20 {
		t.Fatalf("Expected delta if A wins %v, but got %v", 20, delta1)
	}
	delta2 := elo.RatingDelta(A, B, 0)
	if delta2 != -11 {
		t.Fatalf("Expected delta if A loses %v, but got %v", -11, delta2)
	}
	delta3 := elo.RatingDelta(A, B, 0.5)
	if delta3 != 4 {
		t.Fatalf("Expected delta if draw %v, but got %v", 4, delta3)
	}

	rating1 := elo.Rating(A, B, 1)
	if rating1 != 1520 {
		t.Fatalf("Expected rating if A wins %v, but got %v", 1520, rating1)
	}
	rating2 := elo.Rating(A, B, 0)
	if rating2 != 1489 {
		t.Fatalf("Expected rating if A loses %v, but got %v", 1489, rating2)
	}
	rating3 := elo.Rating(A, B, 0.5)
	if rating3 != 1504 {
		t.Fatalf("Expected rating if draw %v, but got %v", 1504, rating3)
	}

	outcomeA1, outcomeB1 := elo.Outcome(A, B, 1) 
	if outcomeA1.Rating != 1520 || outcomeA1.Delta != 20 {
		t.Fatalf("Expected rating %v and delta %v if A wins, but got outcome %v", 1520, 20, outcomeA1)
	}	
	if outcomeB1.Rating != 1580 || outcomeB1.Delta != -20 {
		t.Fatalf("Expected rating %v and delta %v if B loses, but got outcome %v", 1580, -20, outcomeB1)
	}	
}

func TestEloWithFactors(t *testing.T) {
	elo := NewEloWithFactors(40, 800)


	expected := elo.ExpectedScore(A, B)
	if expected != 0.4285368825916186 {
		t.Fatalf("Expected chance %v, but got %v", 0.4285368825916186, expected)
	}

	delta1 := elo.RatingDelta(A, B, 1)
	if delta1 != 22 {
		t.Fatalf("Expected delta if A wins %v, but got %v", 22, delta1)
	}
	delta2 := elo.RatingDelta(A, B, 0)
	if delta2 != -17 {
		t.Fatalf("Expected delta if A loses %v, but got %v", -17, delta2)
	}
	delta3 := elo.RatingDelta(A, B, 0.5)
	if delta3 != 2 {
		t.Fatalf("Expected delta if draw %v, but got %v", 2, delta3)
	}

	rating1 := elo.Rating(A, B, 1)
	if rating1 != 1522 {
		t.Fatalf("Expected rating if A wins %v, but got %v", 1522, rating1)
	}
	rating2 := elo.Rating(A, B, 0)
	if rating2 != 1483 {
		t.Fatalf("Expected rating if A loses %v, but got %v", 1483, rating2)
	}
	rating3 := elo.Rating(A, B, 0.5)
	if rating3 != 1502 {
		t.Fatalf("Expected rating if draw %v, but got %v", 1502, rating3)
	}

	outcomeA1, outcomeB1 := elo.Outcome(A, B, 1) 
	if outcomeA1.Rating != 1522 || outcomeA1.Delta != 22 {
		t.Fatalf("Expected rating %v and delta %v if A wins, but got outcome %v", 1522, 22, outcomeA1)
	}	
	if outcomeB1.Rating != 1578 || outcomeB1.Delta != -22 {
		t.Fatalf("Expected rating %v and delta %v if B loses, but got outcome %v", 1578, -22, outcomeB1)
	}	
}