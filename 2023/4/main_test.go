package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestParseCard(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	cardId, winningNumbers, myNumbers := parseCard(line)

	if cardId != 1 {
		t.Errorf("Expected cardId to be 1, got %d", cardId)
	}
	if len(winningNumbers) != 5 {
		t.Errorf("Expected winningNumbers to be 5, got %d", len(winningNumbers))
	}
	if len(myNumbers) != 8 {
		t.Errorf("Expected myNumbers to be 8, got %d", len(myNumbers))
	}
	if !slices.Equal(winningNumbers, []int{41, 48, 83, 86, 17}) {
		t.Errorf("Expected winningNumbers to be [41, 48, 83, 86, 17], got %d", winningNumbers)
	}
	if !slices.Equal(myNumbers, []int{83, 86, 6, 31, 17, 9, 48, 53}) {
		t.Errorf("Expected myNumbers to be [83, 86, 6, 31, 17, 9, 48, 53], got %d", myNumbers)
	}

	line = "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	cardId, winningNumbers, myNumbers = parseCard(line)

	if cardId != 2 {
		t.Errorf("Expected cardId to be 2, got %d", cardId)
	}
	if len(winningNumbers) != 5 {
		t.Errorf("Expected winningNumbers to be 5, got %d", len(winningNumbers))
	}
	if len(myNumbers) != 8 {
		t.Errorf("Expected myNumbers to be 8, got %d", len(myNumbers))
	}
	if !slices.Equal(winningNumbers, []int{13, 32, 20, 16, 61}) {
		t.Errorf("Expected winningNumbers to be [13, 32, 20, 16, 61], got %d", winningNumbers)
	}
	if !slices.Equal(myNumbers, []int{61, 30, 68, 82, 17, 32, 24, 19}) {
		t.Errorf("Expected myNumbers to be [61, 30, 68, 82, 17, 32, 24, 19], got %d", myNumbers)
	}
}

func TestCountEqualNumbers(t *testing.T) {
	numbers := []int{41, 48, 83, 86, 17}
	myNumbers := []int{83, 86, 6, 31, 17, 9, 48, 53}
	count := countEqualNumbers(numbers, myNumbers)
	if count != 4 {
		t.Errorf("Expected count to be 4, got %d", count)
	}

	numbers = []int{13, 32, 20, 16, 61}
	myNumbers = []int{61, 30, 68, 82, 17, 32, 24, 19}
	count = countEqualNumbers(numbers, myNumbers)
	if count != 2 {
		t.Errorf("Expected count to be 0, got %d", count)
	}
}

func TestCalculatePoints(t *testing.T) {
	card := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	_, winningNumbers, myNumbers := parseCard(card)
	score := countEqualNumbers(winningNumbers, myNumbers)
	points := calculatePoints(score)
	if points != 8 {
		t.Errorf("Expected points to be 8, got %d", points)
	}

	card = "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	_, winningNumbers, myNumbers = parseCard(card)
	score = countEqualNumbers(winningNumbers, myNumbers)
	points = calculatePoints(score)
	if points != 2 {
		t.Errorf("Expected points to be 2, got %d", points)
	}

	card = "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"
	_, winningNumbers, myNumbers = parseCard(card)
	score = countEqualNumbers(winningNumbers, myNumbers)
	points = calculatePoints(score)
	if points != 2 {
		t.Errorf("Expected points to be 2, got %d", points)
	}

	card = "Card 4:  1  2  3  4  5 |  6  7  8  9 10 11 12 13"
	_, winningNumbers, myNumbers = parseCard(card)
	score = countEqualNumbers(winningNumbers, myNumbers)
	points = calculatePoints(score)

	card = "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"
	_, winningNumbers, myNumbers = parseCard(card)
	score = countEqualNumbers(winningNumbers, myNumbers)
	points = calculatePoints(score)
	if points != 0 {
		t.Errorf("Expected points to be 0, got %d", points)
	}

}
