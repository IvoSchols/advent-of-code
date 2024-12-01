package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solution for Day 2 of Advent of Code 2023
func main() {
	fmt.Printf("Part One:%d\n", part_one())
	fmt.Printf("Part Two:%d\n", part_two())
}

func part_one() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	// There are 12 red cubes, 13 green cubes, and 14 blue cubes.
	max_red := 12
	max_green := 13
	max_blue := 14

	sum := 0

	// Read input.txt line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		game_id, bags := getGame(line)
		valid_game := true

		for _, bag := range bags {
			red, blue, green := getColorsInBag(bag)
			if red > max_red || blue > max_blue || green > max_green {
				valid_game = false
				break
			}
		}

		if valid_game {
			sum += game_id
		}

	}

	return sum
}

// Find the minimum number of cubes required to play each game
// Take their power (red * blue * green) and sum them
func part_two() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	sum := 0

	// Read input.txt line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		_, bags := getGame(line)
		min_red, min_blue, min_green := 0, 0, 0

		for _, bag := range bags {
			red, blue, green := getColorsInBag(bag)
			if red > min_red {
				min_red = red
			}
			if blue > min_blue {
				min_blue = blue
			}
			if green > min_green {
				min_green = green
			}
		}
		power := min_red * min_blue * min_green
		sum += power
	}

	return sum

}

func getGame(line string) (int, []string) {
	// Get the game id and the bag string from a line
	// Return the game id and the bag string
	game_string := strings.Split(line, ": ")
	// Get game id by removing "Game " from the beginning of the string
	game_id, _ := strconv.Atoi(game_string[0][5:])
	bags := strings.Split(game_string[1], "; ")
	return game_id, bags
}

func getColorsInBag(bag string) (int, int, int) {
	// Get the number of red, green, and blue cubes in a bag
	// Return the number of red, green, and blue cubes in a bag
	colors := strings.Split(bag, ", ")
	red, blue, green := 0, 0, 0
	for _, color := range colors {
		if strings.Contains(color, "red") {
			// Skip last three characters (" red")
			red, _ = strconv.Atoi(color[:len(color)-4])
		}
		if strings.Contains(color, "blue") {
			blue, _ = strconv.Atoi(color[:len(color)-5])
		}
		if strings.Contains(color, "green") {
			// fmt.Println("green")
			green, _ = strconv.Atoi(color[:len(color)-6])
			// fmt.Println(green)
		}
	}
	return red, blue, green
}
