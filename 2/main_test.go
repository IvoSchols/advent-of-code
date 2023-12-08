package main

import (
	"testing"
)

func TestGetGame(t *testing.T) {
	// Test getGame
	game_id, bags := getGame("Game 1: 12 red, 11 green; 13 green; 14 blue")
	if game_id != 1 {
		t.Errorf("getGame failed, expected %d, got %d", 1, game_id)
	}
	if len(bags) != 3 {
		t.Errorf("getGame failed, expected %d, got %d", 3, len(bags))
	}
	if bags[0] != "12 red, 11 green" {
		t.Errorf("getGame failed, expected %s, got %s", "12 red, 11 green", bags[0])
	}
	if bags[1] != "13 green" {
		t.Errorf("getGame failed, expected %s, got %s", "13 green", bags[1])
	}
	if bags[2] != "14 blue" {
		t.Errorf("getGame failed, expected %s, got %s", "14 blue", bags[2])
	}
}

func TestGetColorsInBag(t *testing.T) {
	// Test getColorsInBag
	red, green, blue := getColorsInBag("12 red, 11 green")
	if red != 12 {
		t.Errorf("getColorsInBag failed, expected %d, got %d", 12, red)
	}
	if green != 11 {
		t.Errorf("getColorsInBag failed, expected %d, got %d", 11, green)
	}
	if blue != 0 {
		t.Errorf("getColorsInBag failed, expected %d, got %d", 0, blue)
	}
}
