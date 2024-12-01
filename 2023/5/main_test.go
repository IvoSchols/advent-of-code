package main

import (
	"testing"
)

func TestIsInMapRange(t *testing.T) {
	// Seed
	start := 47
	length := 27

	// Test the 4 mapping cases and 2 non-mapping cases

	// Mapping case 1: Total mapping
	src := 40
	m_len := 100
	if !isInMapRange(start, length, src, m_len) {
		t.Errorf("Expected isInMapRange to return true, got false")
	}

	// Mapping case 2: Partial mapping left
	src = 40
	m_len = 20
	if !isInMapRange(start, length, src, m_len) {
		t.Errorf("Expected isInMapRange to return true, got false")
	}

	// Mapping case 3: Partial mapping right
	src = 55
	m_len = 48
	if !isInMapRange(start, length, src, m_len) {
		t.Errorf("Expected isInMapRange to return true, got false")
	}

	// Mapping case 4: Partial mapping middle
	src = 55
	m_len = 5
	if !isInMapRange(start, length, src, m_len) {
		t.Errorf("Expected isInMapRange to return true, got false")
	}

	// Non-mapping case 1: Mapping is before the seed
	src = 0
	m_len = 20
	if isInMapRange(start, length, src, m_len) {
		t.Errorf("Expected isInMapRange to return false, got true")
	}

	// Non-mapping case 2: Mapping is after the seed
	src = 100
	m_len = 20
	if isInMapRange(start, length, src, m_len) {
		t.Errorf("Expected isInMapRange to return false, got true")
	}

}

func TestShouldMapCompletely(t *testing.T) {
	// Seed
	start := 47
	length := 27

	// Test the 4 mapping cases
	// Mapping case 1: Total mapping
	src := 40
	m_len := 100
	if !shouldMapCompletely(start, length, src, m_len) {
		t.Errorf("Expected shouldMapCompletely to return true, got false")
	}

	// Mapping case 2: Partial mapping left
	src = 40
	m_len = 20
	if shouldMapCompletely(start, length, src, m_len) {
		t.Errorf("Expected shouldMapCompletely to return false, got true")
	}

	// Mapping case 3: Partial mapping right
	src = 55
	m_len = 48
	if shouldMapCompletely(start, length, src, m_len) {
		t.Errorf("Expected shouldMapCompletely to return false, got true")
	}
	// Mapping case 4: Partial mapping middle
	src = 55
	m_len = 5
	if shouldMapCompletely(start, length, src, m_len) {
		t.Errorf("Expected shouldMapCompletely to return false, got true")
	}
}

func TestShouldMapPartiallyLeft(t *testing.T) {
	// Seed
	start := 47
	length := 27

	// Test the 4 mapping cases
	// Mapping case 1: Total mapping
	src := 40
	m_len := 100
	if shouldMapPartiallyLeft(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyLeft to return false, got true")
	}

	// Mapping case 2: Partial mapping left
	src = 40
	m_len = 20
	if !shouldMapPartiallyLeft(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyLeft to return true, got false")
	}

	// Mapping case 3: Partial mapping right
	src = 55
	m_len = 48
	if shouldMapPartiallyLeft(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyLeft to return false, got true")
	}
	// Mapping case 4: Partial mapping middle
	src = 55
	m_len = 5
	if shouldMapPartiallyLeft(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyLeft to return false, got true")
	}
}

func TestShouldMapPartiallyRight(t *testing.T) {
	// Seed
	start := 47
	length := 27

	// Test the 4 mapping cases
	// Mapping case 1: Total mapping
	src := 40
	m_len := 100
	if shouldMapPartiallyRight(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyRight to return false, got true")
	}

	// Mapping case 2: Partial mapping left
	src = 40
	m_len = 20
	if shouldMapPartiallyRight(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyRight to return false, got true")
	}

	// Mapping case 3: Partial mapping right
	src = 55
	m_len = 48
	if !shouldMapPartiallyRight(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyRight to return true, got false")
	}
	// Mapping case 4: Partial mapping middle
	src = 55
	m_len = 5
	if shouldMapPartiallyRight(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyRight to return false, got true")
	}
}

func TestShouldMapPartiallyMiddle(t *testing.T) {
	// Seed
	start := 47
	length := 27

	// Test the 4 mapping cases
	// Mapping case 1: Total mapping
	src := 40
	m_len := 100
	if shouldMapPartiallyMiddle(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyMiddle to return false, got true")
	}

	// Mapping case 2: Partial mapping left
	src = 40
	m_len = 20
	if shouldMapPartiallyMiddle(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyMiddle to return false, got true")
	}

	// Mapping case 3: Partial mapping right
	src = 55
	m_len = 48
	if shouldMapPartiallyMiddle(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyMiddle to return false, got true")
	}
	// Mapping case 4: Partial mapping middle
	src = 55
	m_len = 5
	if !shouldMapPartiallyMiddle(start, length, src, m_len) {
		t.Errorf("Expected shouldMapPartiallyMiddle to return true, got false")
	}
}

func TestMapSeedComplete(t *testing.T) {
	// Seed
	start := 47
	length := 27
	// Mapping
	dest := 0
	src := 40
	m_len := 100

	mapped := mapSeedCompletely(start, length, src, m_len, dest)

	if mapped[0] != 7 {
		t.Errorf("Expected mapped[0] to be 7, got %d", mapped[0])
	}
	if mapped[1] != 27 {
		t.Errorf("Expected mapped[1] to be 27, got %d", mapped[1])
	}
}

func TestMapPartiallyLeft(t *testing.T) {
	// Seed
	start := 47
	length := 27
	// Mapping
	dest := 0
	src := 40
	m_len := 20

	mapped, unmapped := mapSeedPartiallyLeft(start, length, src, m_len, dest)

	if mapped[0] != 7 {
		t.Errorf("Expected mapped[0] to be 0, got %d", mapped[0])
	}
	if mapped[1] != 13 {
		t.Errorf("Expected mapped[1] to be 13, got %d", mapped[1])
	}
	if unmapped[0] != 60 {
		t.Errorf("Expected unmapped[0] to be 60, got %d", unmapped[0])
	}
	if unmapped[1] != 14 {
		t.Errorf("Expected unmapped[1] to be 14, got %d", unmapped[1])
	}
}

func TestMapSeedPartiallyRight(t *testing.T) {
	// Seed
	start := 47
	length := 27
	// Mapping
	dest := 0
	src := 55
	m_len := 48

	mapped, unmapped := mapSeedPartiallyRight(start, length, src, m_len, dest)

	if mapped[0] != 0 {
		t.Errorf("Expected mapped[0] to be 0, got %d", mapped[0])
	}
	if mapped[1] != 19 {
		t.Errorf("Expected mapped[1] to be 19, got %d", mapped[1])
	}
	if unmapped[0] != 47 {
		t.Errorf("Expected unmapped[0] to be 47, got %d", unmapped[0])
	}
	if unmapped[1] != 8 {
		t.Errorf("Expected unmapped[1] to be 8, got %d", unmapped[1])
	}
}

func TestMapSeedPartiallyMiddle(t *testing.T) {
	// Seed
	start := 47
	length := 27
	// Mapping
	dest := 0
	src := 55
	m_len := 5

	mapped, unmapped_left, unmapped_right := mapSeedPartiallyMiddle(start, length, src, m_len, dest)

	if mapped[0] != 0 {
		t.Errorf("Expected mapped[0] to be 0, got %d", mapped[0])
	}
	if mapped[1] != 5 {
		t.Errorf("Expected mapped[1] to be 5, got %d", mapped[1])
	}
	if unmapped_left[0] != 47 {
		t.Errorf("Expected unmapped_left[0] to be 47, got %d", unmapped_left[0])
	}
	if unmapped_left[1] != 8 {
		t.Errorf("Expected unmapped_left[1] to be 8, got %d", unmapped_left[1])
	}
	if unmapped_right[0] != 60 {
		t.Errorf("Expected unmapped_right[0] to be 60, got %d", unmapped_right[0])
	}
	if unmapped_right[1] != 14 {
		t.Errorf("Expected unmapped_right[1] to be 14, got %d", unmapped_right[1])
	}
}
