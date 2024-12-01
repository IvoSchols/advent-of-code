package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Solution for Day 3 of Advent of Code 2023
func main() {
	fmt.Printf("Part One: %d\n", part_one())
	fmt.Printf("Part Two: %d\n", part_two())
}

func part_one() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	// Build a 2d rune array of the input
	var engine_schematic [][]rune
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		engine_schematic = append(engine_schematic, []rune(line))
	}

	sum := 0

	// Iterate through the engine_schematic
	for i := 0; i < len(engine_schematic); i++ {
		for j := 0; j < len(engine_schematic[i]); j++ {

			// Check if the current space is a digit
			if unicode.IsDigit(engine_schematic[i][j]) {
				part_number, length_part_number := getIntegerInDirection(engine_schematic, i, j, [2]int{0, 1})

				// Get places around the integer to check for a symbol
				places := getPlacesAroundInteger(engine_schematic, i, j, length_part_number)
				// Check if there is a symbol around the integer
				for _, place := range places {
					if isSymbol(engine_schematic, place[0], place[1]) {
						// If there is a symbol, add the integer to the sum
						sum += part_number
						// Skip the length of the integer (we already checked it, row is skipped in outer loop)
						j += length_part_number
						break
					}
				}
			}

		}
	}
	return sum
}

func part_two() int {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	// Build a 2d rune array of the input
	var engine_schematic [][]rune
	var gear_positions [][2]int
	scanner := bufio.NewScanner(inputFile)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		engine_schematic = append(engine_schematic, []rune(line))

		// Find the gear positions
		for j := 0; j < len(engine_schematic[i]); j++ {
			if engine_schematic[i][j] == '*' {
				gear_positions = append(gear_positions, [2]int{i, j})
			}
		}
		i++
	}

	sum := 0

	// Iterate through the gear schematic and find numbers around gears
	for _, gear_position := range gear_positions {
		i, j := gear_position[0], gear_position[1]

		// Get places around the gear to check for a number
		places := getPlacesAroundInteger(engine_schematic, i, j, 1)
		touched_places := make([][2]int, 0)

		gears := make([]int, 0)

		// Check if there is a number around the gear
		for _, place := range places {
			if !isPlaceInTouchedPlaces(place, touched_places) && unicode.IsDigit(engine_schematic[place[0]][place[1]]) {
				// Find root of the number by moving left
				root_i, root_j := place[0], place[1]-1
				for isInBounds(engine_schematic, root_i, root_j) && unicode.IsDigit(engine_schematic[root_i][root_j]) {
					root_j--
				}
				root_j++
				part_number, length_part_number := getIntegerInDirection(engine_schematic, root_i, root_j, [2]int{0, 1})
				gears = append(gears, part_number)

				// Add places of the part number to touched_places
				for k := 0; k < length_part_number; k++ {
					touched_places = append(touched_places, [2]int{root_i, root_j + k})
				}
			}
		}
		// Add the product of the gears to the sum
		if len(gears) < 2 {
			continue
		}
		product := 1
		for _, gear := range gears {
			product *= gear
		}
		sum += product
	}
	return sum
}
func isPlaceInTouchedPlaces(place [2]int, touched_places [][2]int) bool {
	for _, touched_place := range touched_places {
		if touched_place == place {
			return true
		}
	}
	return false
}

// Find all places around integer in direction to check for a symbol
func getPlacesAroundInteger(engine_schematic [][]rune, i int, j int, length int) [][2]int {
	places := make([][2]int, 0)

	// If top-left, left or bottom_left neighbor is within bounds add it to places
	neighbor_i, neighbor_j := i-1, j-1
	if isInBounds(engine_schematic, neighbor_i, neighbor_j) {
		places = append(places, [2]int{neighbor_i, neighbor_j})
	}
	neighbor_i, neighbor_j = i, j-1
	if isInBounds(engine_schematic, neighbor_i, neighbor_j) {
		places = append(places, [2]int{neighbor_i, neighbor_j})
	}
	neighbor_i, neighbor_j = i+1, j-1
	if isInBounds(engine_schematic, neighbor_i, neighbor_j) {
		places = append(places, [2]int{neighbor_i, neighbor_j})
	}

	// If top-right, right or bottom_right neighbor is within bounds add it to places
	neighbor_i, neighbor_j = i-1, j+length
	if isInBounds(engine_schematic, neighbor_i, neighbor_j) {
		places = append(places, [2]int{neighbor_i, neighbor_j})
	}
	neighbor_i, neighbor_j = i, j+length
	if isInBounds(engine_schematic, neighbor_i, neighbor_j) {
		places = append(places, [2]int{neighbor_i, neighbor_j})
	}
	neighbor_i, neighbor_j = i+1, j+length
	if isInBounds(engine_schematic, neighbor_i, neighbor_j) {
		places = append(places, [2]int{neighbor_i, neighbor_j})
	}

	// Check up down for each digit in the integer
	for l := 0; l < length; l++ {
		neighbor_j = j + l

		neighbor_i := i - 1
		if isInBounds(engine_schematic, neighbor_i, neighbor_j) {
			places = append(places, [2]int{neighbor_i, neighbor_j})
		}
		neighbor_i = i + 1
		if isInBounds(engine_schematic, neighbor_i, neighbor_j) {
			places = append(places, [2]int{neighbor_i, neighbor_j})
		}
	}
	return places
}

// Builds a string of the digits in the direction starting from i,j
// Returns the integer value of the string and the length of the string
func getIntegerInDirection(engine_schematic [][]rune, i int, j int, direction [2]int) (int, int) {
	string_digit := string(engine_schematic[i][j])

	neighbor_i := i + direction[0]
	neighbor_j := j + direction[1]

	for isInBounds(engine_schematic, neighbor_i, neighbor_j) {
		if unicode.IsDigit(engine_schematic[neighbor_i][neighbor_j]) {
			string_digit += string(engine_schematic[neighbor_i][neighbor_j])
			neighbor_i += direction[0]
			neighbor_j += direction[1]
		} else {
			break
		}
	}
	digit, _ := strconv.Atoi(string_digit)
	return digit, len(string_digit)
}

func isInBounds(engine_schematic [][]rune, i int, j int) bool {
	return i >= 0 && i < len(engine_schematic) && j >= 0 && j < len(engine_schematic[i])
}

func isSymbol(engine_schematic [][]rune, i int, j int) bool {
	return !(unicode.IsDigit(engine_schematic[i][j]) || engine_schematic[i][j] == '.')
}
